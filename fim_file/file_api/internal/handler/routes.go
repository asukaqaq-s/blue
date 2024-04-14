// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	Admin "fim_server/fim_file/file_api/internal/handler/Admin"
	"fim_server/fim_file/file_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/file/image",
				Handler: ImageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/file/file",
				Handler: FileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/file/:imageName",
				Handler: ImageShowHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/file/files",
					Handler: Admin.FileListHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/file/files",
					Handler: Admin.FileListRemoveHandler(serverCtx),
				},
			}...,
		),
	)
}