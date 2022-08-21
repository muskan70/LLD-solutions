package eviction

import ddl "cacheGo/doublylinkedlist"

type LRUEviction struct {
	IEviction
	lst ddl.DoublyLinkedList
	mp  map[string]*ddl.Node
}

func NewLRUEviction() IEviction {
	return &LRUEviction{lst: ddl.DoublyLinkedList{}, mp: make(map[string]*ddl.Node)}
}

func (l *LRUEviction) KeyAccessed(key string) {
	if val, ok := l.mp[key]; ok {
		l.lst.Remove(val)
	}
	n := l.lst.Add(key)
	l.mp[key] = n
}

func (l *LRUEviction) EvictKey() string {
	n := l.lst.Evict()
	if n != nil {
		delete(l.mp, *n)
		return *n
	}
	return ""
}
