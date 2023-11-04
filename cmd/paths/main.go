package main

import (
	"algo/internal/graph"
	"bpfio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		log.Fatalln("Expected three positional args: GRAPH_PATH ALGORITHM ORIGIN")
	}
	
	path := args[0]
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)
	graph := graph.NewGraphFromStream(lines)

	var pf graph.Search 
	switch algo := args[1]; algo {
	case "DFS":
		pf := graph.NewDepthFirstSearch(graph)
	default:
		fmt.Println("Unknown graph pathfinding algorithm:", algo)
		os.Exit(1)
	}

}
