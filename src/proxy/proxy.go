package proxy

import (
	"connect"
	"logx"
	"net/url"
	"store"
	"time"
)

const (
	originPoolKey = "pool"
	TIME_OUT      = 30
)

type Refresher interface {
	ExtractTestUrl(addresses []string, out chan<- *url.URL, finishSignal chan<- bool)
	Ping(in <-chan *url.URL, out chan<- *connect.TestResult, finishSignal chan<- bool)
	Process(in <-chan *connect.TestResult, finishSignal chan<- bool)
	All() []string
}

type BasePool struct {
	PingBufferSize    int
	ProcessBufferSize int
	PingerCount       int
}

type Pool struct {
	BasePool
	Storage store.SetStringStorer
	Key     string
}

func (pool *BasePool) ExtractTestUrl(addresses []string, out chan<- *url.URL, finishSignal chan<- bool) {
	defer close(out)
	for _, address := range addresses {
		proxy, err := url.Parse("http://" + address)
		if err == nil {
			out <- proxy
		}
	}
	finishSignal <- true
}

func (pool *BasePool) Ping(in <-chan *url.URL, out chan<- *connect.TestResult, finishSignal chan<- bool) {
	defer close(out)
	concurrenceSignal := make(chan bool, pool.PingerCount)
	for proxy := range in {
		concurrenceSignal <- true
		go func(p *url.URL, o chan<- *connect.TestResult, c <-chan bool) {
			o <- connect.Test(p)
			<-c
		}(proxy, out, concurrenceSignal)
	}
	for len(concurrenceSignal) > 0 {
		time.Sleep(100 * time.Millisecond)
	}
	finishSignal <- true
}

func Refresh(refresher Refresher, pingBufferSize, processBufferSize int) {
	urlsBuffer := make(chan *url.URL, pingBufferSize)
	resultsBuffer := make(chan *connect.TestResult, processBufferSize)

	finishExtractSignal := make(chan bool)
	finishPingSignal := make(chan bool)
	finishProcessSignal := make(chan bool)
	go refresher.ExtractTestUrl(refresher.All(), urlsBuffer, finishExtractSignal)
	go refresher.Ping(urlsBuffer, resultsBuffer, finishPingSignal)
	go refresher.Process(resultsBuffer, finishProcessSignal)

	<-finishExtractSignal
	<-finishPingSignal
	<-finishProcessSignal
}

func InitBasePool() BasePool {
	return BasePool{PingBufferSize: 1024, ProcessBufferSize: 1024, PingerCount: 2048}
}

func InitPool() *Pool {
	return &Pool{BasePool: InitBasePool(), Storage: store.RedisSet{}, Key: "PoolKey"}
}

func (pool *BasePool) All() []string {
	return []string{}
}

func (pool *Pool) All() []string {
	return pool.Storage.All(pool.Key)
}

func (pool *Pool) Process(in <-chan *connect.TestResult, finishSignal chan<- bool) {
	for result := range in {
		if result.Time < TIME_OUT && result.Ok {
			pool.Storage.Add(pool.Key, result.Proxy.Host)
		}
	}
	finishSignal <- true
}
