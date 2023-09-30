package sorting

import (
	"golang.org/x/exp/constraints"
)

func SelectionSort(data []constraints.Ordered) {
	n := data.Len()
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data.Less(j, minIdx) {
				minIdx = j
			}
		}
		data.Swap(i, minIdx)
	}
}
