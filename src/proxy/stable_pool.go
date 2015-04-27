package proxy

import (
	"connect"
	"net/url"
	"store"
	"utils"
)

const (
	STABLE_POOL_KEY         = "StablePoolKey"
	MAX_TIME_RECORDS_LENGTH = 100
	STUDY_RECORDS_LENGTH    = 10
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
	times = append([]int{time}, times...)
	if len(times) > MAX_TIME_RECORDS_LENGTH {
		times = times[:MAX_TIME_RECORDS_LENGTH]
	}
	pool.Storage.Set(STABLE_POOL_KEY, proxy, times)
}

func (pool *StablePool) IsEfficient(proxy string) bool {
	times := pool.Storage.Get(STABLE_POOL_KEY, proxy)
	if len(times) < STUDY_RECORDS_LENGTH {
		return false
	}
	for i := len(times) - STUDY_RECORDS_LENGTH; i < len(times); i++ {
		if times[i] < 0 || times[i] > TIME_OUT {
			return false
		}
	}
	return true
}

func (pool *StablePool) ExtractTestUrl(out chan<- *url.URL) {
	defer close(out)
	for _, address := range pool.Storage.Keys(STABLE_POOL_KEY) {
		proxy, err := url.Parse("http://" + address)
		if err == nil {
			out <- proxy
		}
	}
}

func (pool *StablePool) Process(in <-chan *connect.TestResult) {
	for result := range in {
		pool.Add(result.Proxy.Host, utils.EscapeToTime(result.Ok, result.Time))

		if pool.IsEfficient(result.Proxy.Host) {
			efficientPool.Add(result.Proxy.Host)
		} else {
			efficientPool.Remove(result.Proxy.Host)
		}
	}
}
