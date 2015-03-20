package main

import (
	"connect"
	"fmt"
	"net/url"
	"store"
)

func main() {
	rs := &store.RedisSet{}
	test(rs)

	hs := &store.RedisHashArray{}
	test2(hs)

	proxy, err := url.Parse("http://101.251.211.234:80")
	if err != nil {
		panic(proxy)
	}

	time, ok := connect.Test(proxy)
	fmt.Println(time)
	fmt.Println(ok)
}

func test(s store.SetStringStorer) {
	s.Add("s", "a")
	fmt.Println(s.Size("s"))
}

func test2(s store.HashArrayStorer) {
	s.Set("b", "s", []int{1, 2})
	fmt.Println(s.Get("b", "s"))
}
