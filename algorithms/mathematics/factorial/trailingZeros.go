package factorial

// TrailingZeros gets the number of trailing zeros of n!
func TrailingZeros(n int) (count int) {
	for n >= 5 {
		n /= 5
		count += n
	}
	return
}
