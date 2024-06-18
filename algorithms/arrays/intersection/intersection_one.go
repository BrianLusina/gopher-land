package intersection

import "gopherland/pkg/types"

func toSet[T types.Comparable](items []T) map[T]bool {
	set := make(map[T]bool)
	for _, num := range items {
		if _, ok := set[num]; !ok {
			set[num] = true
		}
	}
	return set
}

func setIntersection[T types.Comparable](setOne, setTwo map[T]bool) []T {
	intersection := []T{}
	for key := range setOne {
		if _, ok := setTwo[key]; ok {
			intersection = append(intersection, key)
		}
	}
	return intersection
}

func intersection[T types.Comparable](listOne []T, listTwo []T) []T {
	setOne := toSet(listOne)
	setTwo := toSet(listTwo)

	if len(setOne) < len(setTwo) {
		return setIntersection(setOne, setTwo)
	} else {
		return setIntersection(setTwo, setOne)
	}
}
