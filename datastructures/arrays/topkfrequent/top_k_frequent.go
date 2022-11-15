package topkfrequent

import "container/heap"

// topKFrequent uses a heap to find the top k most frequent elements in the integer slice
func topKFrequent(numbers []int, k int) []int {
	topK := []int{}

	if len(numbers) == 0 {
		return topK
	}

	if k == len(numbers) {
		return numbers
	}

	// Space: O(n)
	frequencyMap := map[int]int{}

	// Time O(n)
	for _, number := range numbers {
		if count, ok := frequencyMap[number]; ok {
			count++
			frequencyMap[number] = count
		} else {
			frequencyMap[number] = 1
		}
	}

	pq := priorityQueue{}

	// Time: O(n)
	for num, count := range frequencyMap {
		// Space O(n * n)
		n := number{value: num, count: count}
		pq = append(pq, n)
	}

	heap.Init(&pq)

	for i := 0; i < k; i++ {
		num := heap.Pop(&pq).(number)
		topK = append(topK, num.value)
	}

	return topK
}

type number struct {
	value int
	count int
}

type priorityQueue []number

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].count >= pq[j].count
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x any) {
	num := x.(number)
	*pq = append(*pq, num)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	l := len(old)
	item := old[l-1]
	// old[l-1] = nil
	*pq = old[:l-1]
	return item
}
