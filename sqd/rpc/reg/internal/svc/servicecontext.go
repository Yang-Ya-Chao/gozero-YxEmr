package svc

import (
	"YxEmr/common/database"
	"YxEmr/sqd/rpc/per/perer"
	"YxEmr/sqd/rpc/reg/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
	Cache   cache.Cache
	Perer   perer.Perer
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.MustSetup(c.Log)
	db := database.Initdb(database.Pubin{
		c.DataSourceName,
	})
	return &ServiceContext{
		Config:  c,
		DbEngin: db,
		Perer:   perer.NewPerer(zrpc.MustNewClient(c.Per)),
	}
}
