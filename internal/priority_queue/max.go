package priority_queue

type MaxPriorityQueue struct {
	data []string
	n    int
}

func NewMaxPriorityQueue(keys []string) *MaxPriorityQueue {
	n := len(keys)
	data := make([]string, n+1)
	for i := 0; i < n; i++ {
		data[i+1] = keys[i]
	}
	pq := &MaxPriorityQueue{data, n}
	for k := n / 2; k >= 1; k-- {
		pq.Sink(k)
	}
	return pq
}

func (pq MaxPriorityQueue) Max() string {
	if len(pq.data) == 0 {
		panic("Queue is empty!")
	}
	return pq.data[1]
}

func (pq *MaxPriorityQueue) Insert(key string) {
	if pq.n == len(pq.data)-1 {
		pq.resize(2 * len(pq.data))
	}
	pq.n++
	pq.data[pq.n] = key
	pq.Swim(pq.n)
}

func (pq *MaxPriorityQueue) DelMax() string {
	if len(pq.data) == 0 {
		panic("Queue is empty!")
	}

	max := pq.data[1]
	pq.swap(1, pq.n)
	pq.n--
	pq.Sink(1)

	if pq.n > 0 && pq.n == (len(pq.data)-1/4) {
		pq.resize(len(pq.data) / 2)
	}

	return max
}

func (pq *MaxPriorityQueue) Sink(node_idx int) {
	for 2*node_idx <= pq.n {
		child_idx := 2 * node_idx
		if child_idx < pq.n && pq.less(child_idx, child_idx+1) {
			child_idx++
		}
		if !pq.less(node_idx, child_idx) {
			break
		}
		pq.swap(node_idx, child_idx)
		node_idx = child_idx
	}
}

func (pq *MaxPriorityQueue) Swim(node_idx int) {
	for node_idx > 1 && pq.less(node_idx/2, node_idx) {
		pq.swap(node_idx/2, node_idx)
		node_idx = node_idx / 2
	}
}

func (pq *MaxPriorityQueue) swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq MaxPriorityQueue) less(i, j int) bool {
	return pq.data[i] < pq.data[j]
}

func (pq *MaxPriorityQueue) resize(capacity int) {
	temp := make([]string, capacity)
	for i := 1; i <= pq.n; i++ {
		temp[i] = pq.data[i]
	}
	pq.data = temp
}
