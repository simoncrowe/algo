package symbol_table

type Interface[K comparable, V any] interface {
	Size() int
	IsEmpty() bool
	Put(K, V)
	Get(K) (V, error)
	Delete(K)
	Contains(K) bool
	Keys() []K
}
