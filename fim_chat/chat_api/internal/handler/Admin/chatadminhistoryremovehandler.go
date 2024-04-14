package handler

import (
	"fim_server/common/response"
	"fim_server/fim_chat/chat_api/internal/logic/Admin"
	"fim_server/fim_chat/chat_api/internal/svc"
	"fim_server/fim_chat/chat_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatAdminHistoryRemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatAdminHistoryRemoveRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewChatAdminHistoryRemoveLogic(r.Context(), svcCtx)
		resp, err := l.ChatAdminHistoryRemove(&req)
		response.Response(r, w, resp, err)

	}
}
