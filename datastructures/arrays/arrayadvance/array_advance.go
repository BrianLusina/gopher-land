package arrayadvance

import "gopherland/math/utils"

func arrayAdvance(a []int) bool {
	furthestReached := 0
	lastIdx := len(a) - 1
	i := 0

	for i <= furthestReached && furthestReached < lastIdx {
		furthestReached = utils.Max(furthestReached, a[i]+i)
		i++
	}

	return furthestReached >= lastIdx
}
