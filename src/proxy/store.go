package proxy

type Sizer interface {
	Size() int
}

type Deleter interface {
	Delete(string, string)
}

type HashArrayStore interface {
	Sizer
	Deleter
	Set(string, string, []int)
	Get(string, string) []int
	All(string) map[string][]int
}

type SetStore interface {
	Sizer
	Deleter
	Add(string, string)
	IsIn(string, string) bool
	Rand(string) string
	All(string) []string
}
