package union_find

type QuickUnion struct {
	parents        []int
	componentCount int
}

func NewQuickUnion(idCount int) *QuickUnion {
	var parents []int = make([]int, idCount)
	for i := 0; i < idCount; i++ {
		parents[i] = i
	}
	return &QuickUnion{parents: parents, componentCount: idCount}
}

func (qu *QuickUnion) Union(p int, q int) {
	rootP := qu.Find(p)
	rootQ := qu.Find(q)
	if rootP == rootQ {
		return
	}
	qu.parents[rootP] = rootQ
	qu.componentCount--
}

func (qu QuickUnion) Find(p int) int {
	var i int = p
	for i != qu.parents[i] {
		i = qu.parents[i]
	}
	return i
}

func (qu QuickUnion) Count() int {
	return qu.componentCount
}
