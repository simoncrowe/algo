package sorting

func SelectionSort(data SortableStrings) {
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
