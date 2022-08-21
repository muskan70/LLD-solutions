package eviction

type IEviction interface {
	KeyAccessed(key string)
	EvictKey()
}
