package handler

import (
	"context"
	"errors"
	"fim_server/common/response"
	"fim_server/fim_file/file_api/internal/logic"
	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"
	"fim_server/fim_file/file_model"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/utils"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileRequest
		if err := httpx.ParseHeaders(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		file, fileHead, err := r.FormFile("file")
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		// 文件上传 用黑名单  exe  php
		nameList := strings.Split(fileHead.Filename, ".") // name.png  1.fengfeng.png 1.fengfeng_xxx.png
		var suffix string
		if len(nameList) > 1 {
			suffix = nameList[len(nameList)-1]
		}

		if utils.InList(svcCtx.Config.BlackList, suffix) {
			response.Response(r, w, nil, errors.New("图片非法"))
			return
		}
		fileData, _ := io.ReadAll(file)
		fileHash := utils.MD5(fileData)

		l := logic.NewFileLogic(r.Context(), svcCtx)
		resp, err := l.File(&req)

		var fileModel file_model.FileModel
		err = svcCtx.DB.Take(&fileModel, "hash = ?", fileHash).Error
		if err == nil {
			resp.Src = fileModel.WebPath()
			logx.Infof("文件 %s hash重复", fileHead.Filename)
			response.Response(r, w, resp, err)
			return
		}

		// 文件重名
		// 在保存文件之前，去读一些文件列表  如果有重名的，算一下它们两个的hash值，一样的就不用写了
		// 它们的hash如果是不一样的，就把最新的这个重命名一下 {old_name}_xxxx.{suffix}

		// 先去拿用户信息
		userResponse, err := svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
			UserIdList: []uint32{uint32(req.UserID)},
		})
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		dirName := fmt.Sprintf("%d_%s", req.UserID, userResponse.UserInfo[uint32(req.UserID)].NickName)

		dirPath := path.Join(svcCtx.Config.UploadDir, "file", dirName)
		_, err = os.ReadDir(dirPath)
		if err != nil {
			os.MkdirAll(dirPath, 0666)
		}

		newFileModel := file_model.FileModel{
			UserID:   req.UserID,
			FileName: fileHead.Filename,
			Size:     fileHead.Size,
			Hash:     fileHash,
			Uid:      uuid.New(),
		}
		newFileModel.Path = path.Join(dirPath, fmt.Sprintf("%s.%s", newFileModel.Uid, suffix))
		err = os.WriteFile(newFileModel.Path, fileData, 0666)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		// 文件信息入库
		err = svcCtx.DB.Create(&newFileModel).Error
		if err != nil {
			logx.Error(err)
			response.Response(r, w, resp, err)
			return
		}

		resp.Src = newFileModel.WebPath()
		response.Response(r, w, resp, err)
	}
}
