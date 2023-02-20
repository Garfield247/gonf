package rediscli

import (
	"fmt"

	"github.com/Garfield247/gonf.git/config"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func SetUp() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisCfg.Host, config.RedisCfg.Port),
		Password: config.RedisCfg.Password,
		DB:       config.RedisCfg.DB,
	})
	RDB = rdb
}
