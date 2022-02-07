package tree

import (
	"fmt"
	"sort"
)

// Record is a record of a tree.
type Record struct {
	ID, Parent int
}

// Node is a node of a tree.
type Node struct {
	ID       int
	Children []*Node
}

type NodeSlice []*Node

func (a NodeSlice) Len() int           { return len(a) }
func (a NodeSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NodeSlice) Less(i, j int) bool { return a[i].ID < a[j].ID }

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	nodes := make([]Node, len(records))
	parents := make([]*Node, len(records))
	seen := make([]bool, len(records))

	for _, record := range records {
		if record.ID >= len(records) {
			return nil, fmt.Errorf("ID %d is too high", record.ID)
		}

		if record.ID != 0 && record.ID <= record.Parent {
			return nil, fmt.Errorf("ID %d has self or later parent %d", record.ID, record.Parent)
		}

		if seen[record.ID] {
			return nil, fmt.Errorf("Record %d is duplicated", record.ID)
		}

		seen[record.ID] = true
		if record.ID != 0 {
			parents[record.ID] = &nodes[record.Parent]
		} else if record.Parent != 0 {
			return nil, fmt.Errorf("root node has non-zero parent %d", record.Parent)
		}
	}

	for i := 1; i < len(nodes); i++ {
		parents[i].Children = append(parents[i].Children, &nodes[i])
	}

	for i, node := range nodes {
		nodes[i].ID = i
		sort.Sort(NodeSlice(node.Children))
	}

	return &nodes[0], nil
}
