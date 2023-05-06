package sorting

import (
	"algo/internal/priority_queue"
)


func HeapSort(data SortableStrings) {
	queue := priority_queue.NewMaxPriorityQueue(data)
	for i := len(data)-1; i >= 0; i-- {
		data[i] = queue.DelMax()
	}
}
