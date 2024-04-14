package handler

import (
	"fim_server/common/response"
	"fim_server/fim_group/group_api/internal/logic/Admin"
	"fim_server/fim_group/group_api/internal/svc"
	"fim_server/fim_group/group_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GroupMessageListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupMessageListRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewGroupMessageListLogic(r.Context(), svcCtx)
		resp, err := l.GroupMessageList(&req)
		response.Response(r, w, resp, err)

	}
}
