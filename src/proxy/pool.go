package proxy

import (
	"connect"
	"net/url"
	"store"
)

const (
	POOL_KEY = "PoolKey"
)

var (
	originPool = InitPool()
)

type Pool struct {
	BasePool
	Storage store.SetStringStorer
}

func InitPool() *Pool {
	return &Pool{BasePool: *InitBasePool(), Storage: store.RedisSet{}}
}

func (pool *Pool) ExtractTestUrl(out chan<- *url.URL) {
	defer close(out)
	for _, address := range pool.Storage.All(POOL_KEY) {
		proxy, err := url.Parse("http://" + address)
		if err == nil {
			out <- proxy
		}
	}
}

func (pool *Pool) Process(in <-chan *connect.TestResult) {
	for result := range in {
		if result.Ok {
			stablePool.Add(result.Proxy.Host, result.Time)
		} else {
			stablePool.Add(result.Proxy.Host, TIME_OUT+1)
		}
	}
}
