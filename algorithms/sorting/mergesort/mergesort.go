package mergesort

import "gopherland/pkg/types"

// MergeSortInPlace merges a list in place using merge sort
func MergeSortInPlace[T types.Number](list []T) {

	mid := len(list) / 2
	leftHalf := list[:mid]
	rightHalf := list[mid:]

	MergeSortInPlace(leftHalf)
	MergeSortInPlace(rightHalf)

	i, j, k := 0, 0, 0

	for i < len(leftHalf) && j < len(rightHalf) {
		if leftHalf[i] < rightHalf[j] {
			list[k] = leftHalf[i]
			i++
		} else {
			list[k] = rightHalf[j]
			j++
		}
		k++
	}

	// Grab any lingering items from left half
	for i < len(leftHalf) {
		list[k] = leftHalf[i]
		i++
		k++
	}

	// Grab any lingering items from right half
	for j < len(rightHalf) {
		list[k] = rightHalf[j]
		j++
		k++
	}
}

// combineLists combines 2 lists into 1 sorted list
func combineLists[T types.Number](listOne, listTwo []T) []T {
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
func MergeSort[T types.Number](thelist []T) []T {
	if len(thelist) <= 1 {
		return thelist
	}

	middleIndex := len(thelist) / 2
	leftHalf := thelist[:middleIndex]
	rightHalf := thelist[middleIndex:]

	leftHalfSorted := MergeSort(leftHalf)
	rightHalfSorted := MergeSort(rightHalf)

	return combineLists(leftHalfSorted, rightHalfSorted)
}
