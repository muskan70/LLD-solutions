package eviction

type IEvictionPolicy interface {
	KeyAccessed(key string)
	EvictKey() string
}
