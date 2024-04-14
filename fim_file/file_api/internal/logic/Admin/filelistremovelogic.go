package Admin

import (
	"context"
	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"
	"fim_server/fim_file/file_model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileListRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListRemoveLogic {
	return &FileListRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileListRemoveLogic) FileListRemove(req *types.FileListRemoveRequest) (resp *types.FileListRemoveResponse, err error) {
	var fileList []file_model.FileModel
	l.svcCtx.DB.Find(&fileList, "id in ?", req.IdList).Delete(&fileList)
	logx.Infof("删除文件个数 %d", len(fileList))
	return
}
