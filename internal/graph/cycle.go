package graph

func HasDirectedCycle(g DirectedGraph) bool {
	marked := make([]bool, g.Verts())
	onStack := make([]bool, g.Verts())
	cycle := false

	// Declare first so that the closure function can recurse
	var dfs func(g DirectedGraph, v int)

	dfs = func(g DirectedGraph, v int) {
		marked[v] = true
		onStack[v] = true

		for _, w := range g.Adj(v) {
			if cycle {
				return
			}
			if !marked[w] {
				dfs(g, w)
			} else if onStack[w] {
				// Back edge: w is on the current recursion stack, so v→w closes a cycle 
				cycle = true
			}
		}
		onStack[v] = false
	}

	for v := 0; v < g.Verts(); v++ {
		if !marked[v] && !cycle {
			dfs(g, v)
		}
	}

	return cycle
}
