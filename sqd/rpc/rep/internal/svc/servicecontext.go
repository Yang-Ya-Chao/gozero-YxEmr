package svc

import (
	"YxEmr/common/database"
	"YxEmr/sqd/rpc/rep/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
	Cache   cache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.MustSetup(c.Log)
	db, ca := database.Initdb(database.Pubin{
		c.DataSourceName,
		c.Cache,
	})
	return &ServiceContext{
		Config:  c,
		DbEngin: db,
		Cache:   ca,
	}
}
