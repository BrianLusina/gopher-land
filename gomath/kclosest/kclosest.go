package kclosest

import "sort"

type Points [][]int

func (p Points) Len() int { return len(p) }

func (p Points) Less(i, j int) bool {
	return p[i][0]*p[i][0]+p[i][1]*p[i][1] < p[j][0]*p[j][0]+p[j][1]*p[j][1]
}

func (p Points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func kClosest(points [][]int, k int) [][]int {
	sort.Sort(Points(points))
	return points[:k]
}
