package storage

type HashmapStorage struct {
	hmap map[string]string
}

func NewHashmapStorage() HashmapStorage {
	return HashmapStorage{hmap: make(map[string]string)}
}

func (hs *HashmapStorage) Add(key, value string) {

}

func (hs *HashmapStorage) Remove(key string) {

}

func (hs *HashmapStorage) Get(key string) string {
	return ""
}
