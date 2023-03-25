package union_find

type QuickFind struct {
	ids            []int
	componentCount int
}

func NewQuickFind(idCount int) *QuickFind {
	var ids []int = make([]int, idCount)
	for i := 0; i < idCount; i++ {
		ids[i] = i
	}
	return &QuickFind{ids: ids, componentCount: idCount}
}

func (qf *QuickFind) Union(p int, q int) {
	pId := qf.ids[p]
	qId := qf.ids[q]

	if pId == qId {
		return
	}

	for i := 0; i < len(qf.ids); i++ {
		if qf.ids[i] == pId {
			qf.ids[i] = qId
		}
	}
	qf.componentCount--
}

func (qf QuickFind) Find(p int) int {
	return qf.ids[p]
}

func (qf QuickFind) Count() int {
	return qf.componentCount
}
