package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if len(slice) <= index || index < 0 {
		return 0, false
	}
	return slice[index], true
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	_, ok := GetItem(slice, index)
	if ok {
		slice[index] = value
		return slice
	}
	return append(slice, value)
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length <= 0 {
		return []int{}
	}
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice = SetItem(slice, i, value)
	}
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	_, ok := GetItem(slice, index)
	if !ok {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
