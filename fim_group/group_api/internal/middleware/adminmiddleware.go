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
		role := r.Header.Get("Role")
		if role != "1" {
			response.Response(r, w, nil, errors.New("权限验证失败"))
			return
		}
		next(w, r)
	}
}
