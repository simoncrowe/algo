package main

import (
	"algo/internal/sorting"
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Expected one positional arg: ALGORITHM")
		os.Exit(1)
	}

	var sortFunc func(sorting.SortableStrings)
	switch algo := args[0]; algo {
	case "Selection":
		sortFunc = sorting.SelectionSort
	case "Insertion":
		sortFunc = sorting.InsertionSort
	case "BinaryInsertion":
		sortFunc = sorting.BinaryInsertionSort
	case "Shell":
		sortFunc = sorting.ShellSort
	case "Merge":
		sortFunc = sorting.MergeSort
	case "MergeBottomUp":
		sortFunc = sorting.MergeSortBottomUp
	case "Quick":
		sortFunc = sorting.QuickSort
	case "QuickThreeWay":
		sortFunc = sorting.QuickSortThreeWay
	case "Heap":
		sortFunc = sorting.HeapSort
	default:
		fmt.Println("Unknown sorting algorithm", algo)
		os.Exit(1)
	}

	data := sorting.SortableStrings{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	sortFunc(data)

	for _, val := range data {
		fmt.Println(val)
	}

}
