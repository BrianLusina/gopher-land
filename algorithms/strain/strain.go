package strain

// Ints Type
type Ints []int

// Lists Type
type Lists [][]int

// Strings type
type Strings []string

// Keep checks if the collection integers meets the predicate f & returns a slice of it in result
func (integers Ints) Keep(f func(int) bool) (result Ints) {
	for _, integer := range integers {
		if f(integer) {
			result = append(result, integer)
		}
	}
	return
}

// Discard checks if the collection does not satisfy the predicate & returns result of collection not meeting predicate
func (integers Ints) Discard(f func(int) bool) (result Ints) {
	for _, i := range integers {
		if !f(i) {
			result = append(result, i)
		}
	}
	return
}

// Keep checks that items in collection meet predicate & returns copy of items that do
func (collection Lists) Keep(predicate func([]int) bool) (filtered Lists) {
	for _, item := range collection {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return
}

// Keep receiver checks that items in strs meet predicate & return result of items that do
func (strs Strings) Keep(predicate func(string) bool) (result Strings) {
	for _, str := range strs {
		if predicate(str) {
			result = append(result, str)
		}
	}

	return
}
