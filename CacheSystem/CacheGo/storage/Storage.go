package storage

type IStorage interface {
	Add(key, value string)
	Remove(key string)
	Get(key string) *string
	PrintAll()
}
