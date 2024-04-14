package logic

import (
	"context"

	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileLogic {
	return &FileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileLogic) File(req *types.FileRequest) (resp *types.FileResponse, err error) {
	resp = new(types.FileResponse)
	return
}
