package graph

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type Graph struct {
	verts int
	edges int
	adj   [][]int
}

func NewGraph(verts int) Graph {
	adj := make([][]int, verts)
	return Graph{verts: verts, edges: 0, adj: adj}
}

func NewGraphFromStream(lines *bufio.Scanner) Graph {
	lines.Scan()
	vertsCount, err := strconv.ParseInt(lines.Text(), 10, 32)
	if err != nil {
		log.Fatalln("Error loading vertices count: ", err)
	}
	verts := int(vertsCount)

	lines.Scan()
	edgesCount, err := strconv.ParseInt(lines.Text(), 10, 32)
	if err != nil {
		log.Fatalln("Error loading edges count: ", err)
	}
	edges := int(edgesCount)

	graph := NewGraph(verts)
	for lines.Scan() {
		edge := strings.Split(lines.Text(), " ")
		origin, err := strconv.ParseInt(edge[0], 10, 32)
		if err != nil {
			log.Fatalln("Error loading first vert of edge: ", err)
		}
		target, err := strconv.ParseInt(edge[1], 10, 32)
		if err != nil {
			log.Fatalln("Error loading second vert of edge: ", err)
		}
		graph.AddEdge(int(origin), int(target))
	}
	if graph.Edges() != edges {
		log.Fatalln("Expected ", edges, " edges. Loaded ", graph.Edges())
	}
	return graph
}

func (g Graph) Verts() int {
	return g.verts
}

func (g Graph) Edges() int {
	return g.edges
}

func (g Graph) validateVertex(v int) {
	if v < 0 || v >= g.Verts() {
		log.Fatalln("Vertex ", v, " is not between 0 and ", g.Verts()-1)
	}
}

func (g *Graph) AddEdge(v int, w int) {
	g.validateVertex(v)
	g.validateVertex(w)
	g.edges++
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
}

func (g Graph) Degree(v int) int {
	g.validateVertex(v)
	return len(g.adj[v])
}

func (g Graph) Adj(v int) []int {
	g.validateVertex(v)
	return g.adj[v]
}
