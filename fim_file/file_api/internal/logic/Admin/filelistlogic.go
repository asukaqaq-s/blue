package Admin

import (
	"context"
	"fim_server/common/list_query"
	"fim_server/common/models"
	"fim_server/fim_file/file_model"

	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListLogic {
	return &FileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileListLogic) FileList(req *types.FileListRequest) (resp *types.FileListResponse, err error) {
	list, count, _ := list_query.ListQuery(l.svcCtx.DB, file_model.FileModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Key:   req.Key,
		},
		Likes: []string{"file_name"},
	})
	resp = new(types.FileListResponse)
	for _, model := range list {
		resp.List = append(resp.List, types.FileListInfoResponse{
			FileName:  model.FileName,
			Size:      model.Size,
			Path:      model.Path,
			CreatedAt: model.CreatedAt.String(),
			ID:        model.ID,
			WebPath:   model.WebPath(),
		})
	}
	resp.Count = count
	return
}
