package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Add zrpc.RpcClientConf
	Del zrpc.RpcClientConf
	Reg zrpc.RpcClientConf
	Cha zrpc.RpcClientConf
	Per zrpc.RpcClientConf
	Rep zrpc.RpcClientConf
}
