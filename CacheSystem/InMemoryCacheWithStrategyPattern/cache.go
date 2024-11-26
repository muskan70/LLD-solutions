package main

import "cacheStrategy/eviction"

type Cache struct {
	size    int
	curSize int
	store   map[string]interface{}
	evict   eviction.IEvictionPolicy
}

func NewCache(sz int, evictionPolicy *eviction.IEvictionPolicy) *Cache {
	return &Cache{size: sz, curSize: 0, store: make(map[string]interface{}), evict: *evictionPolicy}
}

func (c *Cache) Put(key string, val interface{}) {
	if c.curSize == c.size {
		c.curSize--
		ele := c.evict.EvictKey()
		delete(c.store, ele)
	}
	c.curSize++
	c.store[key] = val
	c.evict.KeyAccessed(key)
}

func (c *Cache) Get(key string) interface{} {
	val := c.store[key]
	if val != nil {
		c.evict.KeyAccessed(key)
	}
	return val
}
