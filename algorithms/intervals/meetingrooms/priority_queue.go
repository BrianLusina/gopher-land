package meetingrooms

type priorityQueue []int

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x any) {
	num := x.(int)
	*pq = append(*pq, num)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	l := len(old)
	item := old[l-1]
	*pq = old[:l-1]
	return item
}

func (pq *priorityQueue) Peek() any {
	return (*pq)[0]
}
