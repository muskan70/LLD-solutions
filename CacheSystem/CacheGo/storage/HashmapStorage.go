package storage

import "fmt"

type HashmapStorage struct {
	hmap map[string]string
}

func NewHashmapStorage() IStorage {
	return &HashmapStorage{hmap: make(map[string]string)}
}

func (hs *HashmapStorage) Add(key, value string) {
	hs.hmap[key] = value
}

func (hs *HashmapStorage) Remove(key string) {
	delete(hs.hmap, key)
}

func (hs *HashmapStorage) Get(key string) *string {
	if val, ok := hs.hmap[key]; ok {
		return &val
	}
	return nil
}

func (hs *HashmapStorage) PrintAll() {
	for key, val := range hs.hmap {
		fmt.Println(key, val)
	}
}
