package mergesort

import "gopherland/pkg/types"

// MergeSortInPlace merges a list in place using merge sort
func MergeSortInPlace[T types.Number](collection []T) {

	mid := len(collection) / 2
	leftHalf := collection[:mid]
	rightHalf := collection[mid:]

	MergeSortInPlace(leftHalf)
	MergeSortInPlace(rightHalf)

	i, j, k := 0, 0, 0

	for i < len(leftHalf) && j < len(rightHalf) {
		if leftHalf[i] < rightHalf[j] {
			collection[k] = leftHalf[i]
			i++
		} else {
			collection[k] = rightHalf[j]
			j++
		}
		k++
	}

	// Grab any lingering items from left half
	for i < len(leftHalf) {
		collection[k] = leftHalf[i]
		i++
		k++
	}

	// Grab any lingering items from right half
	for j < len(rightHalf) {
		collection[k] = rightHalf[j]
		j++
		k++
	}
}

// combineLists combines 2 lists into 1 sorted list
func combineLists[T types.Comparable](listOne, listTwo []T) []T {
	listOneIndex, listTwoIndex, mergedList := 0, 0, []T{}

	// Both lists have some items left in them
	for listOneIndex < len(listOne) && listTwoIndex < len(listTwo) {
		if listOne[listOneIndex] <= listTwo[listTwoIndex] {
			mergedList = append(mergedList, listOne[listOneIndex])
			listOneIndex++
		} else {
			mergedList = append(mergedList, listTwo[listTwoIndex])
			listTwoIndex++
		}
	}

	// Grab any lingering items from list one
	for listOneIndex < len(listOne) {
		mergedList = append(mergedList, listOne[listOneIndex])
		listOneIndex++
	}

	// Grab any lingering items from list two
	for listTwoIndex < len(listTwo) {
		mergedList = append(mergedList, listTwo[listTwoIndex])
		listTwoIndex++
	}

	return mergedList
}

// MergeSort merges a list out of place, that is, returns a new list
func MergeSort[T types.Comparable](collection []T) []T {
	if len(collection) <= 1 {
		return collection
	}

	middleIndex := len(collection) / 2
	leftHalf := collection[:middleIndex]
	rightHalf := collection[middleIndex:]

	leftHalfSorted := MergeSort(leftHalf)
	rightHalfSorted := MergeSort(rightHalf)

	return combineLists(leftHalfSorted, rightHalfSorted)
}
