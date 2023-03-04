package database

import (
	"YxEmr/common/cache"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// 只在database中使用
var (
	Db *gorm.DB
)

type Pubin struct {
	Dns string
}

func Initdb(in Pubin) *gorm.DB {
	var err error
	if Db, err = gorm.Open(sqlserver.Open(in.Dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second,  // 慢 SQL 阈值
				LogLevel:      logger.Error, // Log level
				Colorful:      true,         // 禁用彩色打印
			},
		), //常常使用gorm默认日志 并设置日志级别
	}); err != nil {
		panic("dabatase.Initdb err: " + err.Error())
	}

	sqlDB, _ := Db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	//Cache = cache.New(
	//	in.Cache,
	//	syncx.NewSingleFlight(),
	//	cache.NewStat("dc"),
	//	errors.New("CacheNotFound"),
	//	cache.WithExpiry(5*time.Second))
	//return Db, Cache
	cache.Init()
	return Db
}
