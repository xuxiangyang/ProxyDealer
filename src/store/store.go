package store

type SetStringStorer interface {
	All(string) []string
	IsIn(string) bool
	Add(string, string)
	Size(string) int
}

type HashArrayStorer interface {
	Get(string, string) []int
	Set(string, string, []int)
	IsKey(string, string) bool
	Keys(string) []string
	Size(string) int
}
