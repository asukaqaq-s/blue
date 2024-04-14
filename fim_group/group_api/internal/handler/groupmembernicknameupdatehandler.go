package handler

import (
	"fim_server/common/response"
	"fim_server/fim_group/group_api/internal/logic"
	"fim_server/fim_group/group_api/internal/svc"
	"fim_server/fim_group/group_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func groupMemberNicknameUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupMemberNicknameUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewGroupMemberNicknameUpdateLogic(r.Context(), svcCtx)
		resp, err := l.GroupMemberNicknameUpdate(&req)
		response.Response(r, w, resp, err)

	}
}