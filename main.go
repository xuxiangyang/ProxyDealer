package main

import (
	"fmt"
	"proxy"
)

func main() {
	ok, host := proxy.Rand()
	fmt.Println(ok)
	fmt.Println(host)
}
