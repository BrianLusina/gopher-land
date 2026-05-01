package sparsevector

type SparseVector struct {
	dimension       int
	nonZeroElements map[int]int
}

// New initializes a SparseVector from a given slice of integers, storing only the non-zero elements.
func New(nums []int) *SparseVector {
	dimension := len(nums)
	nonZeroElements := make(map[int]int)

	for i, num := range nums {
		if num != 0 {
			nonZeroElements[i] = num
		}
	}

	return &SparseVector{
		dimension:       dimension,
		nonZeroElements: nonZeroElements,
	}
}

// DotProduct computes the dot product of the current SparseVector with another SparseVector.
func (sv *SparseVector) DotProduct(vec *SparseVector) int {
	result := 0
	currentVecElements := sv.nonZeroElements
	otherVecElements := vec.nonZeroElements

	// To optimize the dot product calculation, we iterate over the non-zero elements of the smaller vector.
	if len(currentVecElements) < len(otherVecElements) {
		currentVecElements, otherVecElements = otherVecElements, currentVecElements
	}

	for index, value := range currentVecElements {
		if otherValue, exists := otherVecElements[index]; exists {
			result += value * otherValue
		}
	}

	return result
}
