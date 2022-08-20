package storage

type Storage interface {
	Add(key, value string)
	Remove(key string)
	Get(key string) string
}
