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

func (pq *MaxPriorityQueue) Sink(k int) {
	for 2*k <= pq.n {
		j := 2 * k
		if j < pq.n && pq.less(j, j+1) {
			j++
		}
		if !pq.less(k, j) {
			break
		}
		pq.swap(k, j)
		k = j
	}
}

func (pq *MaxPriorityQueue) Swim(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.swap(k/2, k)
		k = k / 2
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
