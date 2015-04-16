package main

import (
	"proxy"
)

func main() {
	pool := proxy.InitPool()
	proxy.Refresh(pool, pool.PingBufferSize, pool.ProcessBufferSize)
}
