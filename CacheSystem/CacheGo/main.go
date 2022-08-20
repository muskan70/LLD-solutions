package main

import (
	"cacheGo/cache"
	"fmt"
)

func main() {
	c := cache.NewCache()
	fmt.Println(c)
}
