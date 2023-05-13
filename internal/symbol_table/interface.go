package symbol_table

type Interface interface {
	Size() int
	IsEmpty() bool
	Put(string, int)
	Get(string) (int, error)
	Delete(string)
	Contains(string) bool
	Keys() []string
}
