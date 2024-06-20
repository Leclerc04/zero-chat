package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret  string
		AccessExpire  int64
		RefreshExpire int64
	}
	MySQL struct {
		DSN             string
		LogMode         bool
		MaxOpenCons     int
		MaxIdleCons     int
		CreateBatchSize int
	}
	RedisConf redis.RedisConf
}
