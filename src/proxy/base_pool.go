package proxy

import (
	"connect"
	"net/url"
	"time"
)

const (
	TIME_OUT = 30
)

type BasePool struct {
	PingBufferSize    int
	ProcessBufferSize int
	PingerCount       int
}

func InitBasePool() *BasePool {
	return &BasePool{PingBufferSize: 1024, ProcessBufferSize: 1024, PingerCount: 2048}
}

func (pool *BasePool) Ping(in <-chan *url.URL, out chan<- *connect.TestResult) {
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
}
