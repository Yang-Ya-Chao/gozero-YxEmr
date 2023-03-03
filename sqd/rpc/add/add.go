package main

import (
	"YxEmr/sqd/rpc/add/add"
	"YxEmr/sqd/rpc/add/internal/config"
	"YxEmr/sqd/rpc/add/internal/server"
	"YxEmr/sqd/rpc/add/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/add.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		add.RegisterAdderServer(grpcServer, server.NewAdderServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 注册nacos服务
	//sc := []constant.ServerConfig{
	//	*constant.NewServerConfig("127.0.0.1", 8848),
	//}
	//
	//cc := &constant.ClientConfig{
	//	NamespaceId:         "public",
	//	TimeoutMs:           50000,
	//	NotLoadCacheAtStart: true,
	//	LogDir:              "runtime/nacos/log",
	//	CacheDir:            "runtime/nacos/cache",
	//	LogLevel:            "debug",
	//}
	//
	//opts := nacos.NewNacosConfig(c.Name, c.ListenOn, sc, cc)
	//_ = nacos.RegisterService(opts)

	fmt.Printf("Starting %s server at %s...\n", c.Name, c.ListenOn)
	s.Start()
}
