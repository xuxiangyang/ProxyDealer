package proxy

import (
	"http"
)

const (
	GetTestSite  = "http://www.baidu.com"
	PostTestSite = "http://www.baidu.com"
)

var (
	OriginPool       *SetPool
	StablePool       *HashArrayStore
	HeighPerformPool *SetPool
	Set              *SetStore
	Hash             *HashArrayStore
)

func RefreshOriginPool() {

}

func RefreshStablePool() {
	allProxies := OriginPool.All()
	for proxy := range allProxies {
	}
}

func RefreshHeightPerfromPool() {

}

func Rand() {

}

func Feedback(proxyString string, time int) {

}

func validGet(reps http.Response) bool {
}

func validPost(reps http.Response) bool {
}

func init() {
	OriginPool = &SetPool{PoolKey: ORIGIN_POOL_KEY, *Set}
	OriginPool = &SetPool{PoolKey: ORIGIN_POOL_KEY, *Set}
}
