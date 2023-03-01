package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSourceName string // 数据库连接
    Log            logx.LogConf
    DBLog          bool
    Cache      cache.CacheConf    // 手动代码
}
