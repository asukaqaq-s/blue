package handler

import (
	"fim_server/common/response"
	"fim_server/fim_settings/settings_api/internal/logic/Admin"
	"fim_server/fim_settings/settings_api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SettingsInfoUpdadeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Admin.SettingsInfoUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewSettingsInfoUpdadeLogic(r.Context(), svcCtx)
		resp, err := l.SettingsInfoUpdade(&req)
		response.Response(r, w, resp, err)

	}
}
