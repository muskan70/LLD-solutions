package cache

import (
	"cacheGo/eviction"
	"cacheGo/storage"
)

type Cache struct {
	stg  storage.HashmapStorage
	evic eviction.LRUEviction
}

func NewCache() Cache {
	return Cache{stg: storage.NewHashmapStorage(), evic: eviction.NewLRUEviction()}
}
