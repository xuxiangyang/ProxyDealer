package proxy

import (
	"store"
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

func Rand() (ok bool, proxy string) {
	return efficientPool.Storage.Rand(EFFICIENT_POOL_KEY)
}

func FeedBack(ok bool, time int) {

}
