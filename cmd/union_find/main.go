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

	var uf union_find.UnionFind
	switch algo := args[0]; algo {
	case "QuickFind":
		qf := union_find.NewQuickFind(ids_count)
		uf = &qf
	case "QuickUnion":
		qu := union_find.NewQuickUnion(ids_count)
		uf = &qu
	case "WeightedQuickUnion":
		wqu := union_find.NewWeightedQuickUnion(ids_count)
		uf = &wqu
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
		if uf.Find(p) == uf.Find(q) {
			continue
		}
		uf.Union(p, q)
		fmt.Println(p, q)
	}
	fmt.Println(uf.Count(), "components")

}
