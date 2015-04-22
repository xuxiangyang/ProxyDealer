package store

import (
	"encoding/json"
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

type RedisHashArray struct {
}

func (redisSet RedisSet) All(key string) []string {
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

func (redisSet RedisSet) IsIn(key string, value string) bool {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return false
	}
	defer client.Close()

	in, err := client.Cmd("SISMEMBER", key, value).Bool()
	if err != nil {
		logx.Warn(err)
		return false
	}
	return in
}

func (redisSet RedisSet) Add(key, member string) error {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return err
	}
	defer client.Close()

	client.Cmd("SADD", key, member)
	return nil
}

func (redisSet RedisSet) Size(key string) int {
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

func (redisSet RedisSet) Rand(key string) (ok bool, proxy string) {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return false, ""
	}
	defer client.Close()
	return true, client.Cmd("SRANDMEMBER", key).String()
}

func (redisSet RedisSet) Remove(key, value string) {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
	}
	defer client.Close()

	client.Cmd("SREM", key, value)
}

func (redisHashArray RedisHashArray) Get(key, feild string) []int {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return []int{}
	}
	defer client.Close()

	serializedArray, err := client.Cmd("HGET", key, feild).Bytes()
	if err != nil {
		logx.Warn(err)
		return []int{}
	}

	var result []int
	err = json.Unmarshal(serializedArray, &result)
	if err != nil {
		logx.Warn(err)
		return []int{}
	}
	return result
}

func (redisHashArray RedisHashArray) Set(key, field string, data []int) error {
	serializedArray, err := json.Marshal(data)
	if err != nil {
		logx.Warn(err)
		return err
	}

	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return err
	}
	defer client.Close()

	client.Cmd("HSET", key, field, string(serializedArray))
	return nil
}

func (redisHashArray RedisHashArray) IsKey(key, field string) bool {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return false
	}
	defer client.Close()

	exist, err := client.Cmd("HEXISTS", key, field).Bool()
	if err != nil {
		logx.Warn(err)
		return false
	}
	return exist
}

func (redisHashArray RedisHashArray) Keys(key string) []string {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return []string{}
	}
	defer client.Close()

	fields, err := client.Cmd("HKEYS", key).List()
	if err != nil {
		logx.Warn(err)
		return []string{}
	}
	return fields
}

func (redisHashArray RedisHashArray) Size(key string) int {
	client, err := redisPool.Get()
	if err != nil {
		logx.Warn(err)
		return 0
	}
	defer client.Close()

	count, err := client.Cmd("HLEN", key).Int()
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
