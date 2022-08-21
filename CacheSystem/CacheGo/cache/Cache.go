package cache

import (
	"cacheGo/eviction"
	"cacheGo/storage"
)

type Cache struct {
	size    int
	curSize int
	stg     storage.IStorage
	evic    eviction.IEviction
}

func NewCache(sz int, stg storage.IStorage, evic eviction.IEviction) Cache {
	return Cache{size: 7, curSize: 0, stg: stg, evic: evic}
}

func (c *Cache) Put(key, value string) {
	if c.curSize == c.size {
		c.evic.EvictKey()
	}
	c.stg.Add(key, value)
	c.evic.KeyAccessed(key)
}

func (c *Cache) Get(key string) string {
	val := c.stg.Get(key)
	if val != nil {
		c.evic.KeyAccessed(key)
	}
	return *val
}
