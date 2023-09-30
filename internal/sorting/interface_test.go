package sorting

import (
	"golang.org/x/exp/constraints"
	"testing"
)

func TestImplementations(t *testing.T) {
	testCases := []struct {
		Name     string
		SortFunc func([]constraints.Ordered)
	}{
		{"Selection Sort", SelectionSort},
		{"Insertion Sort", InsertionSort},
		{"Binary Insertion Sort", BinaryInsertionSort},
		{"Shell Sort", ShellSort},
		{"Merge Sort", MergeSort},
		{"Bottom-Up Merge Sort", MergeSortBottomUp},
		{"Quick Sort", QuickSort},
		{"Three-Way Quick Sort", QuickSortThreeWay},
		{"Heap Sort", HeapSort},
	}

	for _, tc := range testCases {
		input := constraints.Ordered{"Dog", "Box", "Hat", "Elf", "Bog", "Fan", "Ink", "Cat"}

		tc.SortFunc(input)

		expected := []string{"Bog", "Box", "Cat", "Dog", "Elf", "Fan", "Hat", "Ink"}
		for i := 0; i < 7; i++ {
			if input[i] != expected[i] {
				t.Errorf("%s: value at index  %d is %s, expected %s", tc.Name, i, input[i], expected[i])
			}
		}
	}
}
