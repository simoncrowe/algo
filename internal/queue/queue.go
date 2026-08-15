package queue

import (
	"errors"
)

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	items := []T{}
	return &Queue[T]{items: items}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.items) == 0 {
		var nothing T
		return nothing, errors.New("Queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Data() []T {
	return q.items
}
