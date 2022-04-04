package maxlengthdifference

import (
	"math"
	"sort"
)

type sortBy []string

func (s sortBy) Len() int {
	return len(s)
}

func (s sortBy) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s sortBy) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func MxDifLg(a1, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	sort.Sort(sortBy(a1))
	sort.Sort(sortBy(a2))

	shortest1, longest1 := a1[0], a1[len(a1)-1]
	shortest2, longest2 := a2[0], a2[len(a2)-1]

	short := int(math.Abs(float64(len(longest1) - len(shortest2))))
	long := int(math.Abs(float64(len(longest2) - len(shortest1))))

	if long > short {
		return long
	}
	return short
}
