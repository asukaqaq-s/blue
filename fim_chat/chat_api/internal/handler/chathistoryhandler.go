package handler

import (
	"fim_server/common/response"
	"fim_server/fim_chat/chat_api/internal/logic"
	"fim_server/fim_chat/chat_api/internal/svc"
	"fim_server/fim_chat/chat_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func chatHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatHistoryRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewChatHistoryLogic(r.Context(), svcCtx)
		resp, err := l.ChatHistory(&req)
		response.Response(r, w, resp, err)

	}
}
