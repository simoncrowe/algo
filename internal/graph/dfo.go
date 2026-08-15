package graph

import (
	"algo/internal/queue"
	"slices"
)

type DepthFirstOrder struct {
	marked      []bool
	pre         []int
	post        []int
	preorder    queue.Queue[int]
	postorder   queue.Queue[int]
	preCounter  int
	postCounter int
}

func NewDepthFirstOrder(g DirectedGraph) DepthFirstOrder {
	pre := make([]int, g.Verts())
	post := make([]int, g.Verts())
	postorder := queue.NewQueue[int]()
	preorder := queue.NewQueue[int]()
	marked := make([]bool, g.Verts())
	o := DepthFirstOrder{pre: pre, post: post, preorder: *preorder, postorder: *postorder, preCounter: 0, postCounter: 0, marked: marked}

	for v := 0; v < g.Verts(); v++ {
		if !o.marked[v] {
			o.dfs(g, v)
		}
	}

	return o
}

func (o *DepthFirstOrder) dfs(g DirectedGraph, v int) {
	o.marked[v] = true
	o.pre[v] = o.preCounter
	o.preCounter++
	o.preorder.Enqueue(v)

	for _, w := range g.Adj(v) {
		if !o.marked[w] {
			o.dfs(g, w)
		}
	}

	o.postorder.Enqueue(v)
	o.post[v] = o.postCounter
	o.postCounter++
}

func (o DepthFirstOrder) Pre() []int {
	return o.preorder.Data()
}

func (o DepthFirstOrder) Post() []int {
	return o.postorder.Data()
}

func (o DepthFirstOrder) ReversePost() []int {
	data := o.postorder.Data()
	slices.Reverse(data)
	return data
}

func (o DepthFirstOrder) PreNumber(v int) int {
	return o.pre[v]
}

func (o DepthFirstOrder) PostNumber(v int) int {
	return o.post[v]
}
