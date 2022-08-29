package storage

type IStorage interface {
	Add(key string, value interface{})
	Remove(key string)
	Get(key string) interface{}
	PrintAll()
}
