package graph

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	verts int
	edges int
	adjacent [][]int
}

func NewGraph(verts int) Graph {
	adjacent = make([][]int, verts)
	return Graph{verts: verts, edges: edges, adjacent: adjacent}
}


func NewGraphFromStream(lines bufio.Scanner) Graph{
	lines.Scan()
	vertsCount, err := strconv.ParseInt(lines.Text(), 10, 32)
	if err != nil {
		fmt.Println("Error loading vertices count: ", err)
		os.Exit(1)
	}
	verts := int(vertsCount) 
	
	lines.Scan()
	edgesCount, err := strconv.ParseInt(lines.Text(), 10, 32)
	if err != nil {
		fmt.Println("Error loading edges count: ", err)
		os.Exit(1)
	}
	edges := int(edgesCount) 
	
	graph := NewGraph(edgesCount)
	for scanner.Scan() {
        edge := strings.Split(scanner.Text(), " ")
        origin, err := strconv.ParseInt(edge[0], 10, 32)
        if err != nil {
            fmt.Println("Error loading first vert of edge: ", err)
            os.Exit(1)
        } 
        target, err := strconv.ParseInt(edge[1], 10, 32)
        if err != nil {
            fmt.Println("Error loading second vert of edge: ", err)
            os.Exit(1)
        } 
		graph.AddEdge(int(origin), int(target))
	}
	if graph.Edges() != edges {
		fmt.Println("Expected ", edges, " edges. Loaded ", graph.Edges()
	}
	return graph
}
