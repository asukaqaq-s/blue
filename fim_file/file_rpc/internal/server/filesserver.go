// Code generated by goctl. DO NOT EDIT.
// Source: file_rpc.proto

package server

import (
	"context"

	"fim_server/fim_file/file_rpc/internal/logic"
	"fim_server/fim_file/file_rpc/internal/svc"
	"fim_server/fim_file/file_rpc/types/file_rpc"
)

type FilesServer struct {
	svcCtx *svc.ServiceContext
	file_rpc.UnimplementedFilesServer
}

func NewFilesServer(svcCtx *svc.ServiceContext) *FilesServer {
	return &FilesServer{
		svcCtx: svcCtx,
	}
}

func (s *FilesServer) FileInfo(ctx context.Context, in *file_rpc.FileInfoRequest) (*file_rpc.FileInfoResponse, error) {
	l := logic.NewFileInfoLogic(ctx, s.svcCtx)
	return l.FileInfo(in)
}
