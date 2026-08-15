package union_find

type WeightedQuickUnion struct {
	parents        []int
	treeSizes      []int
	componentCount int
}

func NewWeightedQuickUnion(idCount int) *WeightedQuickUnion {
	var parents []int = make([]int, idCount)
	for i := 0; i < idCount; i++ {
		parents[i] = i
	}
	var sizes []int = make([]int, idCount)
	for i := 0; i < idCount; i++ {
		sizes[i] = 1
	}
	return &WeightedQuickUnion{parents: parents, treeSizes: sizes, componentCount: idCount}
}

func (qu *WeightedQuickUnion) Union(p int, q int) {
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

func (qu WeightedQuickUnion) Find(p int) int {
	var i int = p
	for i != qu.parents[i] {
		i = qu.parents[i]
	}
	return i
}

func (qu WeightedQuickUnion) Count() int {
	return qu.componentCount
}
