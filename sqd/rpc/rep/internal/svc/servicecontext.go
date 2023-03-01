package svc

import "YxEmr/sqd/rpc/rep/internal/config"

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DbEngin: database.Initdb(database.Pubin{
			c.DataSourceName,
			c.DBLog,
		}),
	}
}
