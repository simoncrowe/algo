package graph

import (
	"log"
	"slices"
)

type DepthFirstPaths struct {
	marked []bool
	edgeTo []*int
	source int
}

func NewDepthFirstPaths(g Graph, source int) DepthFirstPaths {
	g.validateVertex(source)

	marked := make([]bool, g.Verts())
	edgeTo := make([]*int, g.Verts())
	paths := DepthFirstPaths{marked: marked, edgeTo: edgeTo, source: source}
	paths.dfs(g, source)
	return paths
}

func (paths *DepthFirstPaths) dfs(g Graph, v int) {
	paths.validateVertex(v)
	paths.marked[v] = true
	for _, w := range g.Adj(v) {
		if !paths.marked[w] {
			paths.edgeTo[w] = &v
			paths.dfs(g, w)
		} 
	}
}

func (paths DepthFirstPaths) HasPathTo(v int) bool {
	paths.validateVertex(v)
	return paths.marked[v]
}

func (paths DepthFirstPaths) PathTo(v int) []int {
	paths.validateVertex(v)

	path := []int{}
	if !paths.HasPathTo(v) {
		return path
	}
		
	for x := v; x != paths.source; x = *paths.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, paths.source)
	slices.Reverse(path)
	return path
}

func (paths DepthFirstPaths) validateVertex(v int) {
	maxVert := len(paths.marked) - 1
	if v < 0 || v > maxVert {
		log.Fatalln("Vertex ", v, " is not between 0 and ", maxVert)
	}
}
