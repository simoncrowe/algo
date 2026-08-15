package graph

import (
	"errors"
)

func TopologicalOrder(g DirectedGraph) ([]int, error) {
	var order []int

	if HasDirectedCycle(g) {
		return order, errors.New("Graph has at least one cycle")
	}

	dfo := NewDepthFirstOrder(g)
	order = dfo.ReversePost()

	return order, nil
}
