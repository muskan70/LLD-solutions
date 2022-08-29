package eviction

import ddl "cacheGo/doublylinkedlist"

type LRUEviction struct {
	lst ddl.DoublyLinkedList
	mp  map[string]*ddl.Node
}

func NewLRUEviction() IEviction {
	return &LRUEviction{lst: ddl.DoublyLinkedList{}, mp: make(map[string]*ddl.Node)}
}

func (l *LRUEviction) KeyAccessed(key string) {
	if val, ok := l.mp[key]; ok {
		l.lst.RemoveNode(val)
	}
	n := l.lst.AddNodeAtEnd(key)
	l.mp[key] = n
}

func (l *LRUEviction) EvictKey() string {
	n := l.lst.RemoveNodeAtHead()
	if n != nil {
		delete(l.mp, *n)
		return *n
	}
	return ""
}
