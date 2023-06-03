package queue

import (
	"errors"
)

type StringQueue struct {
    items []string
}

func NewStringQueue() *StringQueue {
	items := []string{}
	return &StringQueue{items: items}
}

func (q *StringQueue) Enqueue(item string) {
    q.items = append(q.items, item)
}

func (q *StringQueue) Dequeue() (string, error) {
    if len(q.items) == 0 {
        return "", errors.New("Queue is empty")
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item, nil
}

func (q *StringQueue) Size() int {
    return len(q.items)
}


func (q *StringQueue) Data() []string {
    return q.items
}
