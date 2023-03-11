package main

import (
	"algo/internal/union_find"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Expected one positional arg: ALGORITHM")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, err := strconv.ParseInt(scanner.Text(), 10, 32)
	if err != nil {
		fmt.Println("Error loading ids count:", err)
		os.Exit(1)
	}
	ids_count := int(count)
		fmt.Println("Ids count:", ids_count)
		
	var algo union_find.UnionFind
	switch algoName := args[0]; algoName {
	case "QuickFind":
		algo = union_find.NewQuickFind(ids_count)
	}

	for scanner.Scan() {
		ids := strings.Split(scanner.Text(), " ")
		idOne, err := strconv.ParseInt(ids[0], 10, 32)
		if err != nil {
			fmt.Println("Error loading first id:", err)
			os.Exit(1)
		}
		p := int(idOne)
		idTwo, err := strconv.ParseInt(ids[1], 10, 32)
		if err != nil {
			fmt.Println("Error loading second id:", err)
			os.Exit(1)
		} 
		q := int(idTwo)
		algo.Union(p, q)
	}

}
