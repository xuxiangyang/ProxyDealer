package main

import (
	"connect"
	"fmt"
	"net/url"
	"proxy"
)

func main() {
	p, err := url.Parse("http://61.232.6.164:8081")
	if err != nil {
		panic(err)
	}
	a := connect.Test(p)
	fmt.Println(a.Time)
	fmt.Println(a.Ok)

	pool := proxy.InitPool()
	proxy.Refresh(pool, pool.PingBufferSize, pool.ProcessBufferSize)
}
