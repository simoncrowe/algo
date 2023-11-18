package main

import (
	"algo/internal/graph"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	g := graph.NewGraphFromStream(lines)

	parsedSource, err := strconv.ParseInt(args[2], 10, 32)
	if err != nil {
		log.Fatalln("Error parsing origin index as int: ", err)
	}
	source := int(parsedSource)

	var paths graph.IPaths
	switch algo := args[1]; algo {
	case "DFS":
		paths = graph.NewDepthFirstPaths(g, source)
	default:
		log.Fatalln("Unknown graph pathfinding algorithm:", algo)
	}

	for v := 0; v < g.Verts(); v++ {
		if paths.HasPathTo(v) {
			fmt.Printf("%d to %d: ", source, v)
			for x := range paths.PathTo(v) {
				if x == source {
					fmt.Print(x)
				} else {
					fmt.Printf("-%d", x)
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d: not connected\n", source, v)
		}
	}
}
