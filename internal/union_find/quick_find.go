package union_find

type QuickFind struct {
	ids []int
}

func NewQuickFind(idCount int) QuickFind {
	var ids []int = make([]int, idCount)
	for i := 0; i < idCount; i++ {
		ids[i] = i
	}
	return QuickFind{ids: ids}
}

func Union(qf QuickFind, p int, q int) {
	pId := qf.ids[p]
	qId := qf.ids[q]

	for i := 0; i < len(qf.ids); i++ {
		if qf.ids[i] == pId {
			qf.ids[i] = qId
		}
	}
}

func Find(qf QuickFind, p int, q int) bool {
	return qf.ids[p] == qf.ids[q]
}
