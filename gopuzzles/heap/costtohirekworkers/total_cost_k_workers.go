package costtohirekworkers

import (
	"container/heap"
	"gopherland/datastructures/trees/heap/binary"
	"gopherland/math/utils"
)

func totalCost(costs []int, k int, candidates int) int64 {
	// head_workers stores the first k workers
	headWorkers := &binary.Heap[int]{}
	// tail_workers stores at most last k workers without any workers from the first k workers
	tailWorkers := &binary.Heap[int]{}

	heap.Init(headWorkers)
	heap.Init(tailWorkers)

	for i := 0; i < candidates; i++ {
		heap.Push(headWorkers, costs[i])
	}

	for i := utils.Max(candidates, len(costs)-candidates); i < len(costs); i++ {
		heap.Push(tailWorkers, costs[i])
	}

	// next_head stores the index of the next worker to be added to the head_workers in case of a hire from the head_workers
	nextHead := candidates

	// next_tail stores the index of the next worker to be added to the tail_workers in case of a hire from the tail_workers
	nextTail := len(costs) - 1 - candidates

	result := 0

	for i := 0; i < k; i++ {
		if tailWorkers.Len() == 0 || headWorkers.Len() != 0 && headWorkers.Peek() <= tailWorkers.Peek() {
			result += heap.Pop(headWorkers).(int)

			// only refill the queue if there are workers outside the 2 queues
			if nextHead <= nextTail {
				heap.Push(headWorkers, costs[nextHead])
				nextHead++
			}
		} else {
			result += heap.Pop(tailWorkers).(int)
			// only refill the queue if there are workers outside the 2 queues
			if nextHead <= nextTail {
				heap.Push(tailWorkers, costs[nextTail])
				nextTail--
			}
		}
	}

	return int64(result)
}
