package sorting

func ShellSort(data SortableStrings) {
	n := data.Len()

	h := 1
	for h < (n / 3) {
		h = (h * 3) + 1
	}

	for h >= 1 {
		// H-sort the array
		for i := 0; i < n; i++ {
			for j := i; j >= h; j -= h {
				if data.Less(j, j-h) {
					data.Swap(j, j-h)
				} else {
					break
				}
			}
		}
		h /= 3
	}
}
