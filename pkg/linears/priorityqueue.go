package linears

// An Item is something we manage in a priority queue.
type Item struct {
	Value    int
	Priority int
	Idx      int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return len(pq) > 1 && pq[i] != nil && pq[j] != nil && pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Idx = i
	pq[j].Idx = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Idx = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n - 1]
	//old[n - 1] = nil  // avoid memory leak
	item.Idx = -1 // for safety
	*pq = old[0 : n - 1]
	return item
}

