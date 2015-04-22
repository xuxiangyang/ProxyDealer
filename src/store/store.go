package store

type SetStringStorer interface {
	All(key string) []string
	IsIn(key, value string) bool
	Add(key, value string) error
	Size(key string) int
	Rand(key string) (bool, string)
	Remove(key, value string)
}

type HashArrayStorer interface {
	Get(key, field string) []int
	Set(key, field string, data []int) error
	IsKey(key, field string) bool
	Keys(key string) []string
	Size(key string) int
}
