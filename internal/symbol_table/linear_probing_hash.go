package symbol_table

import (
	"errors"
	"github.com/dolthub/maphash"
)

const InitialCapacityLP = 128

type LinearProbingHashTable[K comparable, V any] struct {
	keys   []*K
	values []V
	n      uint64
	m      uint64
	hasher maphash.Hasher[K]
}

func NewLinearProbingHashTable[K comparable, V any]() *LinearProbingHashTable[K, V] {
	return newLinearProbingHashTable[K, V](InitialCapacityLP)
}

func newLinearProbingHashTable[K comparable, V any](size int) *LinearProbingHashTable[K, V] {
	keys := make([]*K, size)
	values := make([]V, size)
	hasher := maphash.NewHasher[K]()
	return &LinearProbingHashTable[K, V]{keys: keys, values: values, n: 0, m: uint64(size), hasher: hasher}
}

func (st LinearProbingHashTable[K, V]) Size() int {
	return int(st.n)
}

func (st LinearProbingHashTable[K, V]) IsEmpty() bool {
	return st.n == 0
}

func (st LinearProbingHashTable[K, V]) Keys() []K {
	keys := make([]K, st.Size())
	i := 0
	for _, key := range st.keys {
		if key != nil {
			keys[i] = *key
			i += 1
		}
	}
	return keys
}

func (st *LinearProbingHashTable[K, V]) Put(key K, value V) {
	if st.n >= st.m/2 {
		st.resize(2 * len(st.keys))
	}
	var i uint64
	for i = st.hash(key); st.keys[i] != nil; i = (i + 1) % st.m {
		if *st.keys[i] == key {
			st.values[i] = value
			return
		}
	}
	st.keys[i] = &key
	st.values[i] = value
	st.n += 1
}
func (st LinearProbingHashTable[K, V]) Get(key K) (V, error) {
	var i uint64
	for i = st.hash(key); st.keys[i] != nil; i = (i + 1) % st.m {
		if *st.keys[i] == key {
			return st.values[i], nil
		}
	}
	var nothing V
	return nothing, errors.New("Key not found")
}

func (st LinearProbingHashTable[K, V]) Contains(key K) bool {
	_, err := st.Get(key)
	return err == nil
}

func (st *LinearProbingHashTable[K, V]) Delete(key K) {
	if !st.Contains(key) {
		return
	}

	i := st.hash(key)
	for key != *st.keys[i] {
		i = (i + 1) % st.m
	}

	st.keys[i] = nil

	// Rehash following keys in the same cluster
	i = (i + 1) % st.m
	for st.keys[i] != nil {
		k := st.keys[i]
		v := st.values[i]
		st.keys[i] = nil
		st.n -= 1
		st.Put(*k, v)
		i = (i + 1) % st.m
	}

	st.n -= 1

	if st.n > 0 && st.n <= st.m/8 {
		st.resize(int(st.m / 2))
	}
}

func (st *LinearProbingHashTable[K, V]) hash(key K) uint64 {
	return st.hasher.Hash(key) % uint64(st.m)
}

func (st *LinearProbingHashTable[K, V]) resize(size int) {
	newST := newLinearProbingHashTable[K, V](size)
	for _, key := range st.Keys() {
		value, _ := st.Get(key)
		newST.Put(key, value)
	}
	*st = *newST
}
