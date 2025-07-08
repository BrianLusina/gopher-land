package sortedsquaredarray

import "gopherland/math/utils"

func SortedSquaredArray(array []int) []int {
	length := len(array)
	if length == 0 {
		return []int{}
	}

	result := make([]int, length)
	left, right := 0, length-1

	for i := length - 1; i >= 0; i-- {
		leftAbs, rightAbs := utils.Abs(array[left]), utils.Abs(array[right])
		if leftAbs > rightAbs {
			square := leftAbs * leftAbs
			result[i] = square
			left++
		} else {
			square := rightAbs * rightAbs
			result[i] = square
			right--
		}
	}

	return result
}

func SortedSquaredArray2(array []int) []int {
	length := len(array)
	if length == 0 {
		return []int{}
	}

	result := make([]int, length)
	left, right := 0, length-1

	for i := length - 1; i >= 0; i-- {
		leftSquare, rightSquare := array[left]*array[left], array[right]*array[right]
		if leftSquare > rightSquare {
			result[i] = leftSquare
			left++
		} else {
			result[i] = rightSquare
			right--
		}
	}

	return result
}
