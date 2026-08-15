package graph

import (
	"testing"
)

// topologicalOrderers are the implementations under test. Adding a new
// implementation here runs it against every case in topologicalOrderCases.
var topologicalOrderers = []struct {
	name  string
	order func(DirectedGraph) ([]int, error)
}{
	{"TopologicalOrderDF", TopologicalOrder},
}

var topologicalOrderCases = []struct {
	name    string
	verts   int
	edges   [][2]int
	wantErr bool
}{
	{
		name:  "tiny DAG",
		verts: 4,
		edges: [][2]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}},
	},
	{
		name:  "DAG with several sources",
		verts: 7,
		edges: [][2]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}, {3, 4}, {5, 4}, {5, 6}, {6, 0}},
	},
	{
		name:  "no edges",
		verts: 3,
	},
	{
		// Two components plus an isolated vertex, so the search has to be
		// restarted from more than one source.
		name:  "disconnected",
		verts: 5,
		edges: [][2]int{{0, 1}, {3, 2}},
	},
	{
		name:  "no vertices",
		verts: 0,
	},
	{
		name:    "self loop",
		verts:   2,
		edges:   [][2]int{{0, 0}, {0, 1}},
		wantErr: true,
	},
	{
		// 0 -> 1 -> 2 -> 3 -> 1, so the cycle does not include the source.
		name:    "cycle below the source",
		verts:   4,
		edges:   [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}},
		wantErr: true,
	},
	{
		// The cycle is in a second component, so the search has to be
		// restarted from a later vertex to find it.
		name:    "cycle in unreachable component",
		verts:   5,
		edges:   [][2]int{{0, 1}, {2, 3}, {3, 4}, {4, 2}},
		wantErr: true,
	},
}

func TestTopologicalOrder(t *testing.T) {
	for _, orderer := range topologicalOrderers {
		t.Run(orderer.name, func(t *testing.T) {
			for _, testCase := range topologicalOrderCases {
				t.Run(testCase.name, func(t *testing.T) {
					graph := NewDirectedGraph(testCase.verts)
					for _, edge := range testCase.edges {
						graph.AddEdge(edge[0], edge[1])
					}

					order, err := orderer.order(graph)

					if testCase.wantErr {
						if err == nil {
							t.Fatalf("Expected an error for a cyclic graph, got order %v", order)
						}
						if len(order) != 0 {
							t.Errorf("Expected no order alongside the error, got %v", order)
						}
						return
					}
					if err != nil {
						t.Fatalf("Expected no error for a DAG, got %v", err)
					}

					position := positions(t, order, graph.Verts())

					// Every edge must point forwards in a topological order.
					for _, edge := range testCase.edges {
						if position[edge[0]] > position[edge[1]] {
							t.Errorf(
								"Expected %d to come before %d in order %v",
								edge[0], edge[1], order,
							)
						}
					}
				})
			}
		})
	}
}

// positions maps each vertex to its index in order, failing the test unless
// order holds every vertex of a graph with the given size exactly once.
func positions(t *testing.T, order []int, verts int) map[int]int {
	t.Helper()

	if len(order) != verts {
		t.Fatalf("Expected %d vertices in the order, got %d", verts, len(order))
	}

	position := make(map[int]int, len(order))
	for i, vert := range order {
		if _, seen := position[vert]; seen {
			t.Errorf("Vertex %d appears more than once in order %v", vert, order)
		}
		position[vert] = i
	}
	for vert := range verts {
		if _, ok := position[vert]; !ok {
			t.Errorf("Expected vertex %d in the order, got %v", vert, order)
		}
	}

	return position
}
