package proxy

import (
	"connect"
	"net/url"
)

type Refresher interface {
	ExtractTestUrl(out chan<- *url.URL)
	Ping(in <-chan *url.URL, out chan<- *connect.TestResult)
	Process(in <-chan *connect.TestResult)
}

func Refresh(refresher Refresher, pingBufferSize, processBufferSize int) {
	urlsBuffer := make(chan *url.URL, pingBufferSize)
	resultsBuffer := make(chan *connect.TestResult, processBufferSize)

	finishExtractSignal := make(chan bool)
	finishPingSignal := make(chan bool)
	finishProcessSignal := make(chan bool)

	go func() {
		refresher.ExtractTestUrl(urlsBuffer)
		finishExtractSignal <- true
	}()
	go func() {
		refresher.Ping(urlsBuffer, resultsBuffer)
		finishPingSignal <- true
	}()
	go func() {
		refresher.Process(resultsBuffer)
		finishProcessSignal <- true
	}()
	<-finishExtractSignal
	<-finishPingSignal
	<-finishProcessSignal
}
