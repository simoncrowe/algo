package sorting

import (
	"golang.org/x/exp/constraints"
)

func MergeSortBottomUp(data []constraints.Ordered) {
	N := data.Len()
	aux := make([]string, N)
	for sz := 1; sz < N; sz = sz + sz {
		for lo := 0; lo < N-sz; lo += sz + sz {
			merge(data, aux, lo, lo+sz-1, min(lo+sz+sz-1, N-1))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
