package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"google.golang.org/grpc/reflection"

	"book/library/rpc/internal/config"
	"book/library/rpc/internal/server"
	"book/library/rpc/internal/svc"
	"book/library/rpc/library"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/library.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		library.RegisterLibraryServer(grpcServer, server.NewLibraryServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
