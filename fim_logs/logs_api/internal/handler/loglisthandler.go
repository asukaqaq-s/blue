package handler

import (
	"fim_server/common/response"
	"fim_server/fim_logs/logs_api/internal/logic"
	"fim_server/fim_logs/logs_api/internal/svc"
	"fim_server/fim_logs/logs_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func logListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogListRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewLogListLogic(r.Context(), svcCtx)
		resp, err := l.LogList(&req)
		response.Response(r, w, resp, err)

	}
}
