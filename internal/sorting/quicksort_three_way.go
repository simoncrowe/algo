package sorting

import (
	"math/rand"
	"time"
)

func QuickSortThreeWay(data SortableStrings) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(data.Len(), data.Swap)
	quickSort(data, 0, data.Len()-1)
}

func quickSortThreeWay(data SortableStrings, lo int, hi int) {
	if hi <= lo {
		return
	}
	
	lt, gt := lo, hi
	pivot := data[lo]
	i := lo + 1
	
	for (i <= gt) {
		if data[i] < pivot {
			data.Swap(lt, i)
			lt ++
			i ++
		} else if data[i] > pivot {
			data.Swap(i, gt)
			gt --
		} else {
			i ++ 
		}
	}

	quickSortThreeWay(data, lo, lt-1)
	quickSortThreeWay(data, gt+1, hi)
}

