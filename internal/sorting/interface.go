package sorting

type SortableStrings []string

func (s SortableStrings) Len() int {
	return len(s)
}

func (s SortableStrings) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortableStrings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
