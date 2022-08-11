package rds

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-redis/redis"
)

var (
	RedisMain *redis.Client
)

func init() {
	initRedis := func(addr string, percore int) *redis.Client {
		opt := redis.Options{
			Addr:         addr,
			DialTimeout:  5 * time.Second,
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 2 * time.Second,
			PoolSize:     runtime.NumCPU() * percore,
			MinIdleConns: 1,
		}
		return redis.NewClient(&opt)
	}

	RedisMain = initRedis("dev-general-01.zenmxapps.com:6384", 1)
	fmt.Println("init redis success")

}
