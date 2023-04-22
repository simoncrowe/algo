package sorting

func MergeSort(data SortableStrings) {
	aux := make([]string, data.Len())
	mergeSort(data, aux, 0, data.Len()-1)
}

func mergeSort(data SortableStrings, aux []string, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSort(data, aux, lo, mid)
	mergeSort(data, aux, mid+1, hi)
	merge(data, aux, lo, mid, hi)
}

func merge(data SortableStrings, aux []string, lo int, mid int, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = data[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			data[k] = aux[j]
			j++
		} else if j > hi {
			data[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			data[k] = aux[j]
			j++
		} else {
			data[k] = aux[i]
			i++
		}
	}
}
