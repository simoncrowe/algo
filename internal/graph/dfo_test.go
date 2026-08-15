package graph

import (
	"slices"
	"testing"
)

// tinyDAG returns:
//
//	0 -> 1 -> 3
//	|         ^
//	v         |
//	2 --------+
func tinyDAG() DirectedGraph {
	graph := NewDirectedGraph(4)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	return graph
}

func TestDepthFirstOrderPre(t *testing.T) {
	order := NewDepthFirstOrder(tinyDAG())

	// 0 is visited first, then its first neighbour 1, then 1's neighbour 3,
	// and finally 0's second neighbour 2.
	expected := []int{0, 1, 3, 2}
	if !slices.Equal(order.Pre(), expected) {
		t.Errorf("Expected preorder %v, got %v", expected, order.Pre())
	}
}

func TestDepthFirstOrderPost(t *testing.T) {
	order := NewDepthFirstOrder(tinyDAG())

	// 3 has no outgoing edges so it is done first; 0 is done last.
	expected := []int{3, 1, 2, 0}
	if !slices.Equal(order.Post(), expected) {
		t.Errorf("Expected postorder %v, got %v", expected, order.Post())
	}
}

func TestDepthFirstOrderReversePost(t *testing.T) {
	order := NewDepthFirstOrder(tinyDAG())

	expected := []int{0, 2, 1, 3}
	if !slices.Equal(order.ReversePost(), expected) {
		t.Errorf("Expected reverse postorder %v, got %v", expected, order.ReversePost())
	}
}

func TestDepthFirstOrderNumbers(t *testing.T) {
	order := NewDepthFirstOrder(tinyDAG())

	// The numbers are the positions of each vertex in the pre/post sequences.
	expectedPre := map[int]int{0: 0, 1: 1, 3: 2, 2: 3}
	for vert, number := range expectedPre {
		if order.PreNumber(vert) != number {
			t.Errorf("Expected pre number of %d to be %d, got %d", vert, number, order.PreNumber(vert))
		}
	}

	expectedPost := map[int]int{3: 0, 1: 1, 2: 2, 0: 3}
	for vert, number := range expectedPost {
		if order.PostNumber(vert) != number {
			t.Errorf("Expected post number of %d to be %d, got %d", vert, number, order.PostNumber(vert))
		}
	}
}

func TestDepthFirstOrderReversePostIsTopological(t *testing.T) {
	graph := NewDirectedGraph(7)
	edges := [][2]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}, {3, 4}, {5, 4}, {5, 6}, {6, 0}}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}

	order := NewDepthFirstOrder(graph)
	reversePost := order.ReversePost()

	if len(reversePost) != graph.Verts() {
		t.Fatalf("Expected %d vertices in reverse postorder, got %d", graph.Verts(), len(reversePost))
	}

	position := make(map[int]int, len(reversePost))
	for i, vert := range reversePost {
		if _, seen := position[vert]; seen {
			t.Errorf("Vertex %d appears more than once in reverse postorder %v", vert, reversePost)
		}
		position[vert] = i
	}

	// In a DAG every edge must point forwards in the reverse postorder.
	for _, edge := range edges {
		if position[edge[0]] > position[edge[1]] {
			t.Errorf(
				"Expected %d to come before %d in reverse postorder %v",
				edge[0], edge[1], reversePost,
			)
		}
	}
}

func TestDepthFirstOrderVisitsDisconnectedVertices(t *testing.T) {
	// Two components plus an isolated vertex, so the search has to be
	// restarted from more than one source.
	graph := NewDirectedGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(3, 2)

	order := NewDepthFirstOrder(graph)

	expectedPre := []int{0, 1, 2, 3, 4}
	pre := slices.Clone(order.Pre())
	slices.Sort(pre)
	if !slices.Equal(pre, expectedPre) {
		t.Errorf("Expected every vertex in the preorder, got %v", order.Pre())
	}

	post := slices.Clone(order.Post())
	slices.Sort(post)
	if !slices.Equal(post, expectedPre) {
		t.Errorf("Expected every vertex in the postorder, got %v", order.Post())
	}

	// 2 is reached from 3, so it must be finished before 3.
	if order.PostNumber(2) > order.PostNumber(3) {
		t.Errorf(
			"Expected 2 (post number %d) to be finished before 3 (post number %d)",
			order.PostNumber(2), order.PostNumber(3),
		)
	}
}

func TestDepthFirstOrderSelfLoop(t *testing.T) {
	graph := NewDirectedGraph(2)
	graph.AddEdge(0, 0)
	graph.AddEdge(0, 1)

	order := NewDepthFirstOrder(graph)

	expected := []int{0, 1}
	if !slices.Equal(order.Pre(), expected) {
		t.Errorf("Expected preorder %v, got %v", expected, order.Pre())
	}
	if !slices.Equal(order.Post(), []int{1, 0}) {
		t.Errorf("Expected postorder %v, got %v", []int{1, 0}, order.Post())
	}
}
