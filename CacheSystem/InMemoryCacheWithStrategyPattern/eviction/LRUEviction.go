package eviction

import ds "cacheStrategy/eviction/dataStructures"

type LRUEviction struct {
	ddl       ds.DoubleLinkedList
	keyPosMap map[string]*ds.Node
}

func NewLRUEviction() IEvictionPolicy {
	return &LRUEviction{
		ddl:       *ds.NewDoubleLinkedList(),
		keyPosMap: make(map[string]*ds.Node),
	}
}

func (lru *LRUEviction) KeyAccessed(key string) {
	if pos, ok := lru.keyPosMap[key]; ok {
		lru.ddl.RemoveNode(pos)
	}
	pos := lru.ddl.AddNodeAtEnd(key)
	lru.keyPosMap[key] = pos
}
func (lru *LRUEviction) EvictKey() string {
	return lru.ddl.RemoveNodeAtHead()
}
