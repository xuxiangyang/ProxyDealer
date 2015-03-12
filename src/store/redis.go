package store

import (
	"github.com/fzzy/radix/extra/pool"
	"logx"
)

const (
	redisAddress     = "localhost:6379"
	redisPoolSize    = 32
	maxRedisWaitTime = 5
)

var (
	redisPool *pool.Pool
)

type RedisSet struct {
}

type RedisHash struct {
}

func (redisSet *RedisSet) All(key string) []string {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return []string{}
	}
	defer client.Close()

	strs, err := client.Cmd("SMEMBERS", key).List()
	if err != nil {
		logx.Warn(err)
		return []string{}
	}
	return strs
}

func (redisSet *RedisSet) IsIn(key string) bool {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return false
	}
	defer client.Close()

	in, err := client.Cmd("SISMEMBER", key).Bool()
	if err != nil {
		logx.Warn(err)
		return false
	}
	return in
}

func (redisSet *RedisSet) Add(key, member string) {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return
	}
	defer client.Close()

	client.Cmd("SADD", key, member)
}

func (redisSet *RedisSet) Size(key string) int {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return 0
	}
	defer client.Close()

	count, err := client.Cmd("SCARD", key).Int()
	if err != nil {
		logx.Warn(err)
		return 0
	}
	return count
}

func init() {
	var err error
	redisPool, err = pool.NewPool("tcp", redisAddress, redisPoolSize)
	if err != nil {
		logx.Error(err)
		panic(err)
	}
}
