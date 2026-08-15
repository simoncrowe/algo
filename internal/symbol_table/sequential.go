package symbol_table

import (
	"errors"
)

type ListNode[K comparable, V any] struct {
	key   K
	value V
	next  *ListNode[K, V]
}

type SequentialSearch[K comparable, V any] struct {
	n     int
	first *ListNode[K, V]
}

func NewSequentialSearch[K comparable, V any]() *SequentialSearch[K, V] {
	return &SequentialSearch[K, V]{n: 0, first: nil}
}

func (st SequentialSearch[K, V]) Size() int {
	return st.n
}

func (st SequentialSearch[K, V]) IsEmpty() bool {
	return st.n == 0
}

func (st SequentialSearch[K, V]) Keys() []K {
	keys := make([]K, st.Size())
	node := st.first
	for i := 0; i < st.Size(); i++ {
		keys[i] = node.key
		node = node.next
	}
	return keys
}

func (st *SequentialSearch[K, V]) Put(key K, value V) {
	node := st.first
	for {
		if node == nil {
			break
		}
		if key == node.key {
			node.value = value
			return
		}
		node = node.next
	}
	st.first = &ListNode[K, V]{key: key, value: value, next: st.first}
	st.n++
}

func (st SequentialSearch[K, V]) Get(key K) (V, error) {
	node := st.first
	for {
		if node == nil {
			break
		}
		if key == node.key {
			return node.value, nil
		}
		node = node.next
	}
	var nothing V
	return nothing, errors.New("Key not found")
}

func (st SequentialSearch[K, V]) Contains(key K) bool {
	_, err := st.Get(key)
	return err == nil
}

func (st *SequentialSearch[K, V]) Delete(key K) {
	if st == nil || st.first == nil {
		return
	}

	if st.first.key == key {
		if st.first.next == nil {
			st.first = nil
		} else {
			st.first = st.first.next
		}
		st.n--
		return
	}

	previous := st.first
	current := st.first.next
	for current != nil {
		if key == current.key {
			if current.next == nil {
				previous.next = nil
			} else {
				previous.next = current.next
			}
			st.n--
		}
		previous = current
		current = current.next
	}
}
