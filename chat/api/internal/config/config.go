package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	MySQL struct {
		DSN             string
		LogMode         bool
		MaxOpenCons     int
		MaxIdleCons     int
		CreateBatchSize int
	}
	Kafka struct {
		Addr      string
		Topic     string
		Partition int
		BatchSize int
	}
	UsercenterRpcConf zrpc.RpcClientConf
	RedisConf         redis.RedisConf
}
