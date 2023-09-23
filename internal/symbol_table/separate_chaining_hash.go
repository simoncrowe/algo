package symbol_table

import "github.com/dolthub/maphash"

const InitialCapacity = 4

type SeparateChainingHashTable[K comparable, V any] struct {
	tables []SequentialSearch[K, V]
	n      int
	hasher maphash.Hasher[K]
}

func NewSeparateChainingHashTable[K comparable, V any]() *SeparateChainingHashTable[K, V] {
	return newSeparateChainingHashTable[K, V](InitialCapacity)
}

func newSeparateChainingHashTable[K comparable, V any](size int) *SeparateChainingHashTable[K, V] {
	tables := make([]SequentialSearch[K, V], size)
	hasher := maphash.NewHasher[K]()
	return &SeparateChainingHashTable[K, V]{tables: tables, n: 0, hasher: hasher}
}

func (st SeparateChainingHashTable[K, V]) Size() int {
	return st.n
}

func (st SeparateChainingHashTable[K, V]) IsEmpty() bool {
	return st.n == 0
}

func (st SeparateChainingHashTable[K, V]) Keys() []K {
	keys := make([]K, st.Size())
	i := 0
	for _, list := range st.tables {
		for _, key := range list.Keys() {
			keys[i] = key
			i += 1
		}
	}
	return keys
}

func (st *SeparateChainingHashTable[K, V]) Put(key K, value V) {
	if st.n >= 10*len(st.tables) {
		st.resize(2 * len(st.tables))
	}

	i := st.hash(key)
	if !st.tables[i].Contains(key) {
		st.n += 1
	}
	st.tables[i].Put(key, value)
}

func (st SeparateChainingHashTable[K, V]) Get(key K) (V, error) {
	i := st.hash(key)
	return st.tables[i].Get(key)
}

func (st SeparateChainingHashTable[K, V]) Contains(key K) bool {
	_, err := st.Get(key)
	return err == nil
}

func (st *SeparateChainingHashTable[K, V]) Delete(key K) {
	i := st.hash(key)
	st.tables[i].Delete(key)
}

func (st *SeparateChainingHashTable[K, V]) hash(key K) uint64 {
	m := uint64(len(st.tables))
	return st.hasher.Hash(key) % m
}

func (st *SeparateChainingHashTable[K, V]) resize(size int) {
	newST := newSeparateChainingHashTable[K, V](size)
	for _, list := range st.tables {
		for _, key := range list.Keys() {
			value, _ := list.Get(key)
			newST.Put(key, value)
		}
	}
	*st = *newST
}
