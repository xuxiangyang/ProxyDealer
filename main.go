package main

import (
	"fmt"
	"store"
)

func main() {
	rs := &store.RedisSet{}
	test(rs)

	hs := &store.RedisHashArray{}
	test2(hs)
}

func test(s store.SetStringStorer) {
	s.Add("s", "a")
	fmt.Println(s.Size("s"))
}

func test2(s store.HashArrayStorer) {
	s.Set("b", "s", []int{1, 2})
	fmt.Println(s.Get("b", "s"))
}
