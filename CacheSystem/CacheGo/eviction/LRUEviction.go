package eviction

type LRUEviction struct {
	lst []string
	mp  map[string]string
}

func NewLRUEviction() LRUEviction {
	return LRUEviction{lst: []string{}, mp: make(map[string]string)}
}
