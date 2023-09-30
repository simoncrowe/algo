package sorting

import (
	"algo/internal/priority_queue"
	"golang.org/x/exp/constraints"
)

func HeapSort(data []constraints.Ordered) {
	queue := priority_queue.NewMaxPriorityQueue(data)
	for i := len(data) - 1; i >= 0; i-- {
		data[i] = queue.DelMax()
	}
}
