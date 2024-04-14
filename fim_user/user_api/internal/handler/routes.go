// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	Admin "fim_server/fim_user/user_api/internal/handler/Admin"
	"fim_server/fim_user/user_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/user/user_info",
				Handler: UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/user/user_info",
				Handler: UserInfoUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/friend_info",
				Handler: friendInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/friends",
				Handler: friendListHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/user/friends",
				Handler: friendNoticeUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/search",
				Handler: searchHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/valid",
				Handler: userValidHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/friends",
				Handler: addFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/valid",
				Handler: userValidListHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/user/valid_status",
				Handler: validStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/user/friends",
				Handler: friendDeleteHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/user/users",
					Handler: Admin.UserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/user/curtail",
					Handler: Admin.UserCurtailHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/user/users",
					Handler: Admin.UserDeleteHandler(serverCtx),
				},
			}...,
		),
	)
}
