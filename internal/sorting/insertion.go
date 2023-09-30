package sorting

import (
	"golang.org/x/exp/constraints"
)

func InsertionSort(data []constraints.Ordered) {
	n := data.Len()
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}
