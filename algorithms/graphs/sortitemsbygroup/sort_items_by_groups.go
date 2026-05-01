package sortitemsbygroup

import (
	"gopherland/datastructures/queues/fifo"
	"gopherland/pkg/utils"
)

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// Assign new group IDs to ungrouped items. -1 means ungrouped
	for i := range n {
		if group[i] == -1 {
			group[i] = m
			m += 1
		}
	}

	// Initialize adjacency lists for item level and group level graphs
	itemGraph := map[int][]int{}
	groupGraph := map[int][]int{}

	// Lists to store in-degree graphs for items and groups
	itemInDegree := utils.MakeArray(n, 0)
	groupInDegree := utils.MakeArray(m, 0)

	// Build dependency graphs
	for current := range n {
		for previous := range beforeItems[current] {
			// Add item-level edge
			itemGraph[previous] = append(itemGraph[previous], current)
			itemInDegree[current]++

			// Add group level edge
			if group[previous] != group[current] {
				groupGraph[group[previous]] = append(groupGraph[group[previous]], group[current])
				groupInDegree[group[current]]++
			}
		}
	}

	var topologicalSort func(totalNodes int, graph map[int][]int, inDegree []int) []int
	topologicalSort = func(totalNodes int, graph map[int][]int, inDegree []int) []int {
		queue := fifo.NewFifoQueue[int]()

		// Add all nodes with 0 in-degree to the queue
		for node := range totalNodes {
			if inDegree[node] == 0 {
				queue.Enqueue(node)
			}
		}

		order := []int{}

		// Process in topological order
		for !queue.IsEmpty() {
			node, err := queue.Dequeue()
			if err != nil {
				panic("Failed to dequeue from queue")
			}
			order = append(order, node)

			// reduce in-degree of all neighbors
			for neighbour := range graph[node] {
				inDegree[neighbour]--
				if inDegree[neighbour] == 0 {
					queue.Enqueue(neighbour)
				}
			}
		}

		if len(order) == totalNodes {
			return order
		}
		return []int{}
	}

	// Topologically sort items and groups
	itemOrder := topologicalSort(n, itemGraph, itemInDegree)
	groupOrder := topologicalSort(m, groupGraph, groupInDegree)

	// If either order is empty, then a cycle exists and no valid order can be found
	if len(itemOrder) == 0 || len(groupOrder) == 0 {
		return []int{}
	}

	// Group Items based on their group ID
	groupToItems := map[int][]int{}
	for item := range itemOrder {
		groupToItems[group[item]] = append(groupToItems[group[item]], item)
	}

	result := []int{}
	for g := range groupOrder {
		result = append(result, groupToItems[g]...)
	}

	return result
}
