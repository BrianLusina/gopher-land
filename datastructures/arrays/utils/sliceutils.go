package utils

// InsertSliceAt inserts a slice in another slice at the given index.
// We make a slice called result with the make function.
// The length of the result should be a sum of len(slice) (the length of the original slice) and len(insertion) (the length of insertion).
// Inserting insertion at an index means that elements of slice from and after index will move len(insertion) indexes forward.
// we first copy all the contents from the 0 index until the index before index to result and store the number of elements copied in at.
// Then we copy insertion into result. The contents from slice insertion must be copied in the result after the contents from slice.
// To handle this, we make the variable n. The statement n += copy(result[n:], insertion) will place the insertion after the elements of slice in result,
// and add the copied elements in n.
// Now, we have to copy the remaining elements from slice if any to result. We’ll copy the elements starting from index from slice to the result starting from n index.
// At last, we’ll return the result
func InsertSliceAt(slice, insertion []interface{}, index int) []interface{} {
	result := make([]interface{}, len(slice)+len(insertion))
	n := copy(result, slice[:index])
	n += copy(result[n:], insertion)
	copy(result[n:], slice[index:])
	return result
}
