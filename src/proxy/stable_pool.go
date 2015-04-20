package proxy

import (
	"store"
)

const (
	STABLE_POOL_KEY = "StablePoolKey"
)

var (
	stablePool = InitStablePool()
)

type StablePool struct {
	BasePool
	Storage store.HashArrayStorer
}

func InitStablePool() *StablePool {
	return &StablePool{BasePool: *InitBasePool(), Storage: store.RedisHashArray{}}
}

func (pool *StablePool) Add(proxy string, time int) {
	times := pool.Storage.Get(STABLE_POOL_KEY, proxy)
	times = append(times, time)
	pool.Storage.Set(STABLE_POOL_KEY, proxy, times)
}
