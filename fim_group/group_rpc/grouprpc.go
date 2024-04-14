package main

import (
	"fim_server/common/zrpc_interceptor"
	"flag"
	"fmt"

	"fim_server/fim_group/group_rpc/internal/config"
	"fim_server/fim_group/group_rpc/internal/server"
	"fim_server/fim_group/group_rpc/internal/svc"
	"fim_server/fim_group/group_rpc/types/group_rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/grouprpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		group_rpc.RegisterGroupsServer(grpcServer, server.NewGroupsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(zrpc_interceptor.ServerUnaryInterceptor)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
