package main

import (
	"cacheStrategy/eviction"
	"fmt"
)

func main() {
	cache1 := NewCache(3, eviction.NewLRUEviction())
	cache1.Put("muskan", "mangla")
	cache1.Put("vipul", 12)
	cache1.Put("muskan", 28)
	cache1.Put("sammar", "got a job")
	cache1.Put("shivam", "doing CA")
	fmt.Println(cache1.Get("muskan"))

	cache2 := NewCache(3, eviction.NewLFUEviction())
	cache2.Put("muskan", "mangla")
	cache2.Put("vipul", 12)
	cache2.Put("muskan", 28)
	cache2.Put("sammar", "got a job")
	fmt.Println(cache2.Get("vipul"))
	cache2.Put("shivam", "doing CA")
	fmt.Println(cache2.Get("muskan"))
}
