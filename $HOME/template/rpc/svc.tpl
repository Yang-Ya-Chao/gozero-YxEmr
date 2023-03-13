package svc

import {{.imports}}

type ServiceContext struct {
	Config config.Config
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
