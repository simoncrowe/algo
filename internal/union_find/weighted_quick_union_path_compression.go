package union_find

type WeightedQuickUnionPathComp struct {
	parents        []int
	treeSizes      []int
	componentCount int
}

func NewWeightedQuickUnionPathComp(idCount int) *WeightedQuickUnionPathComp {
	var parents []int = make([]int, idCount)
	for i := 0; i < idCount; i++ {
		parents[i] = i
	}
	var sizes []int = make([]int, idCount)
	for i := 0; i < idCount; i++ {
		sizes[i] = 1
	}
	return &WeightedQuickUnionPathComp{parents: parents, treeSizes: sizes, componentCount: idCount}
}

func (qu *WeightedQuickUnionPathComp) Union(p int, q int) {
	rootP := qu.Find(p)
	rootQ := qu.Find(q)
	if rootP == rootQ {
		return
	}

	if qu.treeSizes[rootP] < qu.treeSizes[rootQ] {
		qu.parents[rootP] = rootQ
		qu.treeSizes[rootQ] += qu.treeSizes[rootP]
	} else {
		qu.parents[rootQ] = rootP
		qu.treeSizes[rootP] += qu.treeSizes[rootQ]
	}
	qu.componentCount--
}

func (qu WeightedQuickUnionPathComp) Find(p int) int {
	var root int = p
	for root != qu.parents[root] {
		root = qu.parents[root]
	}
	var i int = p
	for i != root {
		newi := qu.parents[i]
		qu.parents[i] = root
		i = newi
	}

	return root
}

func (qu WeightedQuickUnionPathComp) Count() int {
	return qu.componentCount
}
