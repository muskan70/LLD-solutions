package storage

import (
	"fmt"
)

type HashmapStorage struct {
	hmap map[string]interface{}
}

func NewHashmapStorage() IStorage {
	return &HashmapStorage{hmap: make(map[string]interface{})}
}

func (hs *HashmapStorage) Add(key string, value interface{}) {
	hs.hmap[key] = value
}

func (hs *HashmapStorage) Remove(key string) {
	delete(hs.hmap, key)
}

func (hs *HashmapStorage) Get(key string) interface{} {
	if val, ok := hs.hmap[key]; ok {
		return val
	}
	return nil
}

func (hs *HashmapStorage) PrintAll() {
	for key, val := range hs.hmap {
		fmt.Println(key, val)
	}
}
