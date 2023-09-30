package sorting

import (
	"golang.org/x/exp/constraints"
)

func (s []constraints.Ordered) Len() int {
	return len(s)
}

func (s []constraints.Ordered) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s []constraints.Ordered) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
