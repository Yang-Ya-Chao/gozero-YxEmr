package svc

import (
	"YxEmr/common/database"
	"YxEmr/sqd/rpc/cha/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.MustSetup(c.Log)
	db := database.Initdb(database.Pubin{
		c.DataSourceName,
	})
	return &ServiceContext{
		Config:  c,
		DbEngin: db,
	}
}
