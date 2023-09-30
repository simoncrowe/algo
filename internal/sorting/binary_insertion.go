package sorting

import (
	"golang.org/x/exp/constraints"
)

func BinaryInsertionSort(data []constraints.Ordered) {
	n := data.Len()
	for i := 1; i < n; i++ {

		// Binary serch for where to insert item i
		v := data[i]
		lo, hi := 0, i
		for lo < hi {
			mid := lo + (hi-lo)/2
			if data.Less(i, mid) {
				hi = mid
			} else {
				lo = mid + 1
			}
		}

		// Shift values to right of where item i is inserted
		for j := i; j > lo; j-- {
			data[j] = data[j-1]
		}
		data[lo] = v
	}
}
