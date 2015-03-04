package proxy

import (
	"github.com/fzzy/radix/extra/pool"
	"github.com/fzzy/radix/redis"
	"time"
)

const (
	maxPickCount = 5
)

var (
	redisPool *pool.Pool
)

func pickARedis() *redis.Client {
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
