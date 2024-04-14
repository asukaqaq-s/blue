package handler

import (
	"fim_server/common/response"
	"fim_server/fim_file/file_api/internal/logic/Admin"
	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileListRemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileListRemoveRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewFileListRemoveLogic(r.Context(), svcCtx)
		resp, err := l.FileListRemove(&req)
		response.Response(r, w, resp, err)

	}
}
