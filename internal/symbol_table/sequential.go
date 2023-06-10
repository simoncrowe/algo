package symbol_table

import (
	"errors"
)

type ListNode struct {
	key   string
	value int
	next  *ListNode
}

type SequentialSearch struct {
	n     int
	first *ListNode
}

func NewSequentialSearch() *SequentialSearch {
	return &SequentialSearch{n: 0, first: nil}
}

func (st SequentialSearch) Size() int {
	return st.n
}

func (st SequentialSearch) IsEmpty() bool {
	return st.n == 0
}

func (st SequentialSearch) Keys() []string {
	keys := make([]string, st.Size())
	node := st.first
	for i := 0; i < st.Size(); i++ {
		keys[i] = node.key
		node = node.next
	}
	return keys
}

func (st *SequentialSearch) Put(key string, value int) {
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
	st.first = &ListNode{key: key, value: value, next: st.first}
	st.n++
}

func (st SequentialSearch) Get(key string) (int, error) {
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
	return 0, errors.New("Key not found")
}

func (st SequentialSearch) Contains(key string) bool {
	_, err := st.Get(key)
	return err == nil
}

func (st *SequentialSearch) Delete(key string) {
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
