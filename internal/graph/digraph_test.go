package graph

import (
	"bufio"
	"strings"
	"testing"
)

func TestDirectedGraphAddEdge(t *testing.T) {
	graph := NewDirectedGraph(5)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)

	if graph.Verts() != 5 {
		t.Errorf("Expected 5 vertices, got %d", graph.Verts())
	}
	if graph.Edges() != 3 {
		t.Errorf("Expected 3 edges, got %d", graph.Edges())
	}
}

func TestDirectedGraphAdj(t *testing.T) {
	graph := NewDirectedGraph(5)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)

	adj := graph.Adj(0)
	if len(adj) != 2 {
		t.Errorf("Expected 0 to be adjacent to 2 vertices, got %d", len(adj))
	}
	if adj[0] != 1 || adj[1] != 2 {
		t.Errorf("Expected 0 to be adjacent to 1 and 2, got %d and %d", adj[0], adj[1])
	}

	// Edges only point one way, so 1 has no adjacent vertices
	if len(graph.Adj(1)) != 0 {
		t.Errorf("Expected 1 to have no adjacent vertices, got %d", len(graph.Adj(1)))
	}
}

func TestDirectedGraphDegree(t *testing.T) {
	graph := NewDirectedGraph(5)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 3)

	if graph.Degree(0) != 3 {
		t.Errorf("Expected the out degree of 0 to be 3, got %d", graph.Degree(0))
	}
	if graph.Degree(3) != 0 {
		t.Errorf("Expected the out degree of 3 to be 0, got %d", graph.Degree(3))
	}
}

func TestDirectedGraphInDegree(t *testing.T) {
	graph := NewDirectedGraph(5)

	graph.AddEdge(0, 3)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)

	if graph.InDegree(3) != 3 {
		t.Errorf("Expected the in degree of 3 to be 3, got %d", graph.InDegree(3))
	}
	if graph.InDegree(0) != 0 {
		t.Errorf("Expected the in degree of 0 to be 0, got %d", graph.InDegree(0))
	}
}

func TestNewDirectedGraphFromStream(t *testing.T) {
	data := "4\n4\n0 1\n0 2\n1 3\n2 3\n"
	lines := bufio.NewScanner(strings.NewReader(data))

	graph := NewDirectedGraphFromStream(lines)

	if graph.Verts() != 4 {
		t.Errorf("Expected 4 vertices, got %d", graph.Verts())
	}
	if graph.Edges() != 4 {
		t.Errorf("Expected 4 edges, got %d", graph.Edges())
	}
	if graph.Degree(0) != 2 {
		t.Errorf("Expected the out degree of 0 to be 2, got %d", graph.Degree(0))
	}
	if graph.InDegree(3) != 2 {
		t.Errorf("Expected the in degree of 3 to be 2, got %d", graph.InDegree(3))
	}
}
