package middleware

import (
	"errors"
	"fim_server/common/response"
	"net/http"
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
	return &AdminMiddleware{}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 只有角色为1的才能调用
		role := r.Header.Get("Role")
		if role != "1" {
			response.Response(r, w, nil, errors.New("角色鉴权失败"))
			return
		}
		next(w, r)
	}
}
