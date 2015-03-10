package proxy

import (
	"github.com/fzzy/radix/extra/pool"
	"github.com/fzzy/radix/redis"
	"time"
)

const (
	maxPickCount  = 5
	redisPoolSize = 32
)

var (
	redisPool *pool.Pool
)

func dialRedis() *redis.Client {
	var redisClient *redis.Client
	var err error = nil
	tryCount := 0
	for err == nil {
		redisClient, err = redisPool.Get()
		if tryCount < maxPickCount {
			Sleep(1)
		} else {

		}
	}
	return redisClient
}

func init() {
	redisPool, err = pool.NewPool("tcp", "localhost:6378", redisPoolSize)
	if err != nil {
		panic(err)
	}
}
