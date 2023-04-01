package sorting

import (
	"testing"
)

func TestSelection(t *testing.T) {
	testCases := []struct {
		Name     string
		SortFunc func(SortableStrings)
	}{
		{"Selection Sort", SelectionSort},
		{"Insertion Sort", InsertionSort},
		{"Binary Insertion Sort", BinaryInsertionSort},
		{"Shell Sort", ShellSort},
	}

	for _, tc := range testCases {
		input := SortableStrings{"Dog", "Box", "Hat", "Elf", "Bog", "Fan", "Ink", "Cat"}

		tc.SortFunc(input)

		expected := []string{"Bog", "Box", "Cat", "Dog", "Elf", "Fan", "Hat", "Ink"}
		for i := 0; i < 7; i++ {
			if input[i] != expected[i] {
				t.Errorf("%s: value at index  %d is %s, expected %s", tc.Name, i, input[i], expected[i])
			}
		}
	}
}
