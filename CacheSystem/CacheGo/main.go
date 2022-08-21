package main

import (
	"cacheGo/cache"
	"cacheGo/eviction"
	"cacheGo/storage"
	"fmt"
)

func main() {
	c := cache.NewCache(3, storage.NewHashmapStorage(), eviction.NewLRUEviction())
	fmt.Println(c)
	c.Put("hello", "muskan")
	fmt.Println("---result:", c.Get("hello"), "---")
	c.Put("hi", "vipul")
	c.Put("good", "papa")
	c.Put("moti", "mummy")
	fmt.Println("---result:", c.Get("good"), "---")
	c.Put("hello", "muskan")
	fmt.Println("---result:", c.Get("good"), "---")
	fmt.Println("---result:", c.Get("moti"), "---")
	c.Put("hi", "vipul")
}
