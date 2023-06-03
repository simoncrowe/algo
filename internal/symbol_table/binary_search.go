package symbol_table

import (
	"errors"
)

type BinarySearch struct {
	n      int
	keys   []string
	values []int
}

func NewBinarySearch() *BinarySearch {
	keys := make([]string, 4096)
	values := make([]int, 4096)
	return &BinarySearch{keys: keys, values: values, n: 0}
}

func (st *BinarySearch) resize(capacity int) {
	newKeys := make([]string, capacity)
	newValues := make([]int, capacity)
	for i := 0; i < st.n; i++ {
		newKeys[i] = st.keys[i]
		newValues[i] = st.values[i]
	}
	st.values = newValues
	st.keys = newKeys
}

func (st BinarySearch) Size() int {
	return st.n
}

func (st BinarySearch) IsEmpty() bool {
	return st.n == 0
}

func (st BinarySearch) Keys() []string {
	return st.keys[:st.n]
}

func (st BinarySearch) Get(key string) (int, error) {
	if st.IsEmpty() {
		return 0, errors.New("Key not found")
	}
	i := st.search(key)
	if i < st.n && st.keys[i] == key {
		return st.values[i], nil
	}

	return 0, errors.New("Key not found")
}

func (st BinarySearch) search(key string) int {
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

func (st *BinarySearch) Put(key string, value int) {
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

func (st BinarySearch) Contains(key string) bool {
	_, err := st.Get(key)
	return err == nil
}

func (st *BinarySearch) Delete(key string) {
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
