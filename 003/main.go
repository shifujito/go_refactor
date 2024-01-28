package main

import (
	"fmt"
	"sample/003/redis"
)

func main() {
	err := redis.Store("foo", "bar")
	fmt.Println(err)
}
