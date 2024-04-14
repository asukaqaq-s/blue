package handler

import (
	"errors"
	"fim_server/common/response"
	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"
	"fim_server/fim_file/file_model"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"os"
)

func ImageShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		var fileModel file_model.FileModel
		err := svcCtx.DB.Take(&fileModel, "uid = ?", req.ImageName).Error
		if err != nil {
			response.Response(r, w, nil, errors.New("文件不存在"))
			return
		}

		byteData, err := os.ReadFile(fileModel.Path)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		w.Write(byteData)
	}
}
