package grains

import "errors"

// Total returns the total number of grains on the board.
func Total() uint64 {
	// the number of grains added up to any point on the board is simply
	// two to the power of that number - 1 (because maths)
	return 1<<64 - 1
}

// Square returns the number of grains on the given square.
// where the first square has 1 and every subsequent square doubles the number.
func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, errors.New("invalid input. Input should be between 1 and 64")
	}
	// a left shift is equivalent to multiplying by 2.  So we need to left
	// shift by num-1 times to get the number of grains on that square
	return 1 << (uint64(i) - 1), nil
}
