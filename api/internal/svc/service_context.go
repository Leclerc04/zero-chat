package svc

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"zero-chat/api/internal/config"
	"zero-chat/api/internal/middleware"
	"zero-chat/api/internal/model"
)

type ServiceContext struct {
	Config       config.Config
	Authority    rest.Middleware
	Redis        *redis.Redis
	DB           *gorm.DB
	UserModel    model.UserModel
	MessageModel model.MessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := Init(c)
	rds := redis.MustNewRedis(c.RedisConf)
	return &ServiceContext{
		Config:       c,
		Authority:    middleware.NewAuthorityMiddleware().Handle,
		DB:           db,
		Redis:        rds,
		UserModel:    model.NewUserModel(db),
		MessageModel: model.NewMessageModel(db),
	}
}

func Init(c config.Config) (db *gorm.DB) {
	var (
		sqlDB *sql.DB
		err   error
	)
	mysqlConf := mysql.Config{DSN: c.MySQL.DSN}

	gormConfig := configLog(c.MySQL.LogMode)
	if db, err = gorm.Open(mysql.New(mysqlConf), gormConfig); err != nil {
		log.Fatal("opens database failed: ", err)
	}
	if sqlDB, err = db.DB(); err != nil {
		log.Fatal("db.db() failed: ", err)
	}

	sqlDB.SetMaxIdleConns(c.MySQL.MaxIdleCons)
	sqlDB.SetMaxOpenConns(c.MySQL.MaxOpenCons)
	return
}

// configLog 根据配置决定是否开启日志
func configLog(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	}
	return
}
