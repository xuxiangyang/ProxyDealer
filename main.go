package main

import (
	"fmt"
	"store"
)

func main() {
	rs := &store.RedisSet{}
	test(rs)
}

func test(s store.SetStringStorer) {
	s.Add("s", "a")
	fmt.Println(s.Size("s"))
}
