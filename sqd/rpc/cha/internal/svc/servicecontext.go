package svc

import (
	"YxEmr/common/database"
	"YxEmr/sqd/rpc/cha/internal/config"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.MustSetup(c.Log)
	return &ServiceContext{
		Config: c,
		DbEngin: database.Initdb(database.Pubin{
			c.DataSourceName,
			c.RedisHost,
			c.DBLog,
		}),
	}
}
