package eviction

import ds "cacheStrategy/eviction/dataStructures"

type LFUEviction struct {
	freqListMap map[int]*ds.DoubleLinkedList
	keyFreqMap  map[string]int
	minFreq     int
}

func NewLFUEviction() IEvictionPolicy {
	return &LFUEviction{
		freqListMap: make(map[int]*ds.DoubleLinkedList),
		keyFreqMap:  make(map[string]int),
		minFreq:     0,
	}
}

func (lfu *LFUEviction) KeyAccessed(key string) {
	freq, ok := lfu.keyFreqMap[key]
	if ok {
		lfu.freqListMap[freq].RemoveKey(key)
	}
	lfu.keyFreqMap[key] = freq + 1
	if lfu.minFreq == 0 {
		lfu.minFreq = 1
	} else {
		lfu.minFreq = min(lfu.keyFreqMap[key], lfu.minFreq)
	}
	_, ok = lfu.freqListMap[freq+1]
	if !ok {
		lfu.freqListMap[freq+1] = ds.NewDoubleLinkedList()
	}
	lfu.freqListMap[freq+1].AddNodeAtEnd(key)
}

func (lfu *LFUEviction) EvictKey() string {
	key := lfu.freqListMap[lfu.minFreq].RemoveNodeAtHead()
	delete(lfu.keyFreqMap, key)
	if len(lfu.keyFreqMap) == 0 {
		lfu.minFreq = 0
	} else {
		for lfu.freqListMap[lfu.minFreq].IsEmpty() {
			lfu.minFreq++
		}
	}
	return key
}
