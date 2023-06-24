package main

import (
	"algo/internal/symbol_table"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Expected one positional arg: ALGORITHM")
		os.Exit(1)
	}

	var table symbol_table.Interface[string, int]
	switch algo := args[0]; algo {
	case "SequentialSearch":
		table = symbol_table.NewSequentialSearch[string, int]()
	case "BinarySearch":
		table = symbol_table.NewBinarySearch[string, int]()
	case "BinarySearchTree":
		table = symbol_table.NewBinarySearchTree[string, int]()
	case "RedBlackBST":
		table = symbol_table.NewRedBlackBST[string, int]()
	default:
		fmt.Println("Unknown symbol table search algorithm", algo)
		os.Exit(1)
	}
	distinct, words := 0, 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		key := scanner.Text()
		for _, word := range strings.Split(key, " ") {
			words++
			if table.Contains(word) {
				count, err := table.Get(word)
				if err != nil {
					panic("Key contained but not found")
				}
				table.Put(word, count+1)
			} else {
				table.Put(word, 1)
				distinct++
			}
		}
	}

	var max string
	max = ""
	table.Put(max, 0)
	for _, key := range table.Keys() {
		currentMax, err := table.Get(max)
		if err != nil {
			panic("Key expected but not found")
		}
		keyMax, err := table.Get(key)
		if err != nil {
			panic("Key expected but not found")
		}
		if keyMax > currentMax {
			max = key
		}
	}
	fmt.Println("Key count: ", table.Size())
	maxCount, err := table.Get(max)
	if err != nil {
		msg := fmt.Sprintf("Expected key not found: '%s'", max)
		panic(msg)
	}
	fmt.Printf("Most frequent word: '%s' (%d)\n", max, maxCount)
	fmt.Printf("Total words: %d\n", words)
	fmt.Printf("Total distinct: %d\n", distinct)
}
