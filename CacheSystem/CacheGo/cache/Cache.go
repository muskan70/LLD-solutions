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
	return Cache{size: sz, curSize: 0, stg: stg, evic: evic}
}

func (c *Cache) Put(key string, value interface{}) {
	if c.curSize == c.size {
		c.curSize--
		m := c.evic.EvictKey()
		c.stg.Remove(m)
	}
	if c.stg.Add(key, value) {
		c.curSize++
	}
	c.evic.KeyAccessed(key)
	c.stg.PrintAll()
}

func (c *Cache) Get(key string) interface{} {
	val := c.stg.Get(key)
	if val != nil {
		c.evic.KeyAccessed(key)
		return val
	}
	return nil
}
