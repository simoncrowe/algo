package sorting

import (
	"math/rand"
	"time"
)

func QuickSort(data SortableStrings) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(data.Len(), data.Swap)
	quickSort(data, 0, data.Len()-1)
}

func quickSort(data SortableStrings, lo int, hi int) {
	if hi <= lo {
		return
	}
	j := partition(data, lo, hi)
	quickSort(data, lo, j-1)
	quickSort(data, j+1, hi)
}

func partition(data SortableStrings, lo int, hi int) int {
	i, j := lo, hi+1
	pivot := data[lo]
	for {
		for {
			i++
			if data[i] >= pivot || i == hi {
				break
			}
		}
		for {
			j--
			if data[j] <= pivot || j == lo {
				break
			}
		}

		if i >= j {
			// Indices cross
			break
		}
		data.Swap(i, j)
	}

	data.Swap(lo, j)
	return j
}
