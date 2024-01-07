package pqueue

type Item[K any] struct {
	Value    K   // The value of the item; arbitrary.
	Priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A priorityQueue implements heap.Interface and holds Items.
type PriorityQueue[K any] []*Item[K]

func (pq PriorityQueue[K]) Len() int { return len(pq) }

func (pq PriorityQueue[K]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue[K]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[K]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[K])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[K]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
