package eviction

type Eviction interface {
	KeyAccessed(key string)
	EvictKey()
}
