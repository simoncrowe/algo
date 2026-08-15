package graph

import (
	"testing"
)

func TestHasDirectedCycleEmptyGraph(t *testing.T) {
	graph := NewDirectedGraph(3)

	if HasDirectedCycle(graph) {
		t.Error("Expected a graph with no edges to have no cycle")
	}
}

func TestHasDirectedCycleDAG(t *testing.T) {
	if HasDirectedCycle(tinyDAG()) {
		t.Error("Expected the tiny DAG to have no cycle")
	}
}

func TestHasDirectedCycleSelfLoop(t *testing.T) {
	graph := NewDirectedGraph(2)
	graph.AddEdge(0, 0)
	graph.AddEdge(0, 1)

	if !HasDirectedCycle(graph) {
		t.Error("Expected a self loop to count as a cycle")
	}
}

func TestHasDirectedCycleTwoVertexCycle(t *testing.T) {
	graph := NewDirectedGraph(2)
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 0)

	if !HasDirectedCycle(graph) {
		t.Error("Expected 0 -> 1 -> 0 to be a cycle")
	}
}

func TestHasDirectedCycleLongerCycle(t *testing.T) {
	// 0 -> 1 -> 2 -> 3 -> 1, so the cycle does not include the source.
	graph := NewDirectedGraph(4)
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 1)

	if !HasDirectedCycle(graph) {
		t.Error("Expected 1 -> 2 -> 3 -> 1 to be a cycle")
	}
}

func TestHasDirectedCycleIgnoresCrossEdges(t *testing.T) {
	// 3 is reached twice, but the second time it is already finished rather
	// than on the stack, so this is not a cycle.
	//
	//	0 -> 1 -> 3
	//	|    ^    |
	//	v    |    v
	//	2 ---+    4
	graph := NewDirectedGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 1)
	graph.AddEdge(3, 4)

	if HasDirectedCycle(graph) {
		t.Error("Expected an edge to an already finished vertex not to count as a cycle")
	}
}

func TestHasDirectedCycleInUnreachableComponent(t *testing.T) {
	// The cycle is in a second component, so the search has to be restarted
	// from a later vertex to find it.
	graph := NewDirectedGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 2)

	if !HasDirectedCycle(graph) {
		t.Error("Expected a cycle unreachable from vertex 0 to be found")
	}
}

func TestHasDirectedCycleDisconnectedAcyclic(t *testing.T) {
	// Two components plus an isolated vertex, none of them cyclic.
	graph := NewDirectedGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(3, 2)

	if HasDirectedCycle(graph) {
		t.Error("Expected a disconnected acyclic graph to have no cycle")
	}
}
