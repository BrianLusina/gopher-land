package binarygap

import "strconv"

// BinaryGap determines the longest binary gap in an integer
func BinaryGap(n int) int {
	n64 := int64(n)

	binaryRepr := strconv.FormatInt(n64, 2)

	result := 0
	one := "1"
	indexArr := []int{}
	gaps := []int{}

	for idx, binary := range binaryRepr {
		if string(binary) == one {
			indexArr = append(indexArr, idx)
		}
	}

	for i := len(indexArr); i > 0; i-- {
		if i-2 >= 0 {
			gaps = append(gaps, indexArr[i-1]-indexArr[i-2]-1)
		}
	}

	for _, gap := range gaps {
		if gap > result {
			result = gap
		}
	}

	return result
}
