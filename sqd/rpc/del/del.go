package main

import (
	"YxEmr/sqd/rpc/del/del"
	"YxEmr/sqd/rpc/del/internal/config"
	"YxEmr/sqd/rpc/del/internal/server"
	"YxEmr/sqd/rpc/del/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/del.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		del.RegisterDelerServer(grpcServer, server.NewDelerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting %s server at %s...\n", c.Name, c.ListenOn)
	s.Start()
}
