package eviction

import ds "cacheStrategy/eviction/dataStructures"

type LFUEviction struct {
	pq         ds.PriorityQueue
	keyFreqMap map[string]int
}

func NewLFUEviction() IEvictionPolicy {
	return &LFUEviction{}
}

func (lfu *LFUEviction) KeyAccessed(key string) {
	lfu.keyFreqMap[key]++
}
func (lfu *LFUEviction) EvictKey() string {
	return ""
}
