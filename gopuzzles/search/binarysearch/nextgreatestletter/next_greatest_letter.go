package nextgreatestletter

func nextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1

	for left <= right {
		mid := (left + right) / 2

		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left == len(letters) {
		return letters[0]
	} else {
		return letters[left]
	}
}
