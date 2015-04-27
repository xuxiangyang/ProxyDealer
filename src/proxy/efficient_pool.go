package proxy

import (
	"store"
	"utils"
)

const (
	EFFICIENT_POOL_KEY = "EfficientPoolKey"
)

var (
	efficientPool = InitEfficientPool()
)

type EfficientPool struct {
	BasePool
	Storage store.SetStringStorer
}

func InitEfficientPool() *EfficientPool {
	return &EfficientPool{BasePool: *InitBasePool(), Storage: store.RedisSet{}}
}

func (pool *EfficientPool) Add(proxy string) {
	pool.Storage.Add(EFFICIENT_POOL_KEY, proxy)
}

func (pool *EfficientPool) Remove(proxy string) {
	pool.Storage.Remove(EFFICIENT_POOL_KEY, proxy)
}

func Rand() (ok bool, proxy string) {
	return efficientPool.Storage.Rand(EFFICIENT_POOL_KEY)
}

func Feedback(proxy string, ok bool, time int) {
	stablePool.Add(proxy, utils.EscapeToTime(ok, time))

	if !ok {
		efficientPool.Remove(proxy)
	}
}
