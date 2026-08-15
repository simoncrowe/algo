package graph

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type DirectedGraph struct {
	verts    int
	edges    int
	adj      [][]int
	inDegree []int
}

func NewDirectedGraph(verts int) DirectedGraph {
	adj := make([][]int, verts)
	inDegree := make([]int, verts)
	return DirectedGraph{verts: verts, edges: 0, adj: adj, inDegree: inDegree}
}

func NewDirectedGraphFromStream(lines *bufio.Scanner) DirectedGraph {
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

	graph := NewDirectedGraph(verts)
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

func (g DirectedGraph) Verts() int {
	return g.verts
}

func (g DirectedGraph) Edges() int {
	return g.edges
}

func (g DirectedGraph) validateVertex(v int) {
	if v < 0 || v >= g.Verts() {
		log.Fatalln("Vertex ", v, " is not between 0 and ", g.Verts()-1)
	}
}

func (g *DirectedGraph) AddEdge(v int, w int) {
	g.validateVertex(v)
	g.validateVertex(w)
	g.adj[v] = append(g.adj[v], w)
	g.edges++
	g.inDegree[w]++
}

func (g DirectedGraph) Degree(v int) int {
	g.validateVertex(v)
	return len(g.adj[v])
}

func (g DirectedGraph) Adj(v int) []int {
	g.validateVertex(v)
	return g.adj[v]
}

func (g DirectedGraph) InDegree(v int) int {
	g.validateVertex(v)
	return g.inDegree[v]
}
