package symbol_table

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type BinarySearch[K constraints.Ordered, V any] struct {
	n      int
	keys   []K
	values []V
}

func NewBinarySearch[K constraints.Ordered, V any]() *BinarySearch[K, V] {
	keys := make([]K, 4096)
	values := make([]V, 4096)
	return &BinarySearch[K, V]{keys: keys, values: values, n: 0}
}

func (st *BinarySearch[K, V]) resize(capacity int) {
	newKeys := make([]K, capacity)
	newValues := make([]V, capacity)
	for i := 0; i < st.n; i++ {
		newKeys[i] = st.keys[i]
		newValues[i] = st.values[i]
	}
	st.values = newValues
	st.keys = newKeys
}

func (st BinarySearch[K, V]) Size() int {
	return st.n
}

func (st BinarySearch[K, V]) IsEmpty() bool {
	return st.n == 0
}

func (st BinarySearch[K, V]) Keys() []K {
	return st.keys[:st.n]
}

func (st BinarySearch[K, V]) Get(key K) (V, error) {
	if st.IsEmpty() {
		var nothing V
		return nothing, errors.New("Key not found")
	}
	i := st.search(key)
	if i < st.n && st.keys[i] == key {
		return st.values[i], nil
	}

	var nothing V
	return nothing, errors.New("Key not found")
}

func (st BinarySearch[K, V]) search(key K) int {
	lo, hi := 0, st.n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < st.keys[mid] {
			hi = mid - 1
		} else if key > st.keys[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

func (st *BinarySearch[K, V]) Put(key K, value V) {
	i := st.search(key)

	if i < st.n && st.keys[i] == key {
		st.values[i] = value
		return
	}

	if st.n == len(st.keys) {
		st.resize(2 * len(st.keys))
	}

	for j := st.n; j > i; j-- {
		st.keys[j] = st.keys[j-1]
		st.values[j] = st.values[j-1]
	}
	st.keys[i] = key
	st.values[i] = value
	st.n++
}

func (st BinarySearch[K, V]) Contains(key K) bool {
	_, err := st.Get(key)
	return err == nil
}

func (st *BinarySearch[K, V]) Delete(key K) {
	i := st.search(key)

	if i < st.n && st.keys[i] != key {
		return
	}

	if st.n == len(st.keys) {
		st.resize(2 * len(st.keys))
	}

	for j := st.n; j < i; j++ {
		st.keys[j] = st.keys[j+1]
		st.values[j] = st.values[j+1]
	}
	st.n--
}
