package storage

type IStorage interface {
	Add(key string, value interface{}) bool
	Remove(key string)
	Get(key string) interface{}
	PrintAll()
}
