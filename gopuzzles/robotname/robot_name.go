package robotname

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
)

var robotNameSpec = []string{
	letters,
	letters,
	digits,
	digits,
	digits,
}

// digitIncrements stores the increment value of each digit position for later
// use during the encoding step.
var digitIncrements = make([]int, len(robotNameSpec))

// MaxNames is the maximum number of names that can be generated based on the
// defined robotNameSpec.
var MaxNames int

// Keep track of unused integer IDs. These are intended to be used to
// deterministically generate new robot names, and then be removed from this
// slice.
var availableIDs []int

// Robot defines properties of a robot.
type Robot struct {
	name string
}

func init() {
	x := 1

	specLength := len(robotNameSpec)

	for i := specLength - 1; i >= 0; i-- {
		digitIncrements[i] = x
		x *= len(robotNameSpec[i])
	}
	MaxNames = x

	availableIDs = make([]int, MaxNames)
	for i := 0; i < MaxNames; i++ {
		availableIDs[i] = i
	}
}

func randomInt(x int) int {
	num := big.NewInt(int64(x))
	rnd, _ := rand.Int(rand.Reader, num)
	return int(rnd.Int64())
}

// encodeIntToString encodes a given integer into a multi-base string according
// to the given list of strings, spec, where each string represents a valid symbol
// for its digit position in order of increasing value. The spec describes digits in
// this way in order of most significant digit to least significant digit.
// index 0 is most significant digit, index 1 is second most significant digit, etc.
func encodeIntToString(id int, spec []string) string {
	remainder := id
	var position, minVal int
	var result string

	for i, digitChars := range spec {
		minVal = digitIncrements[i]

		// The character for this digit position is determined by the number of
		// times that the minimum/increment value of that digit can evenly
		// divide into the remainder from the last digit's calculations (or the
		// original ID in the case of the most significant digit).
		position = remainder / minVal
		result += string(digitChars[position])

		// Calculate the remainder to be used for the next digit.
		remainder %= minVal
	}

	return result
}

// generateUniqueId returns random unused ID and returns an error if none are left.
func generateUniqueId() (int, error) {
	numAvailable := len(availableIDs)
	if numAvailable == 0 {
		return 0, fmt.Errorf("no more unique IDs available")
	}

	x := randomInt(numAvailable)
	uniqueId := availableIDs[x]

	// Delete the ith element by replacing it with the last element in the
	// array and cutting off the original last element.
	availableIDs[x] = availableIDs[numAvailable-1]
	availableIDs = availableIDs[:numAvailable-1]

	return uniqueId, nil
}

// generateUniqueName generates a unique name for a robot
func generateUniqueName() (string, error) {
	uniqueId, err := generateUniqueId()
	if err != nil {
		return "", err
	}
	return encodeIntToString(uniqueId, robotNameSpec), nil
}

// Name returns the name of the robot.
func (r *Robot) Name() (string, error) {
	var err error
	if r.name == "" {
		r.name, err = generateUniqueName()
		if err != nil {
			return "", err
		}
	}
	return r.name, nil
}

// Reset resets a robot
func (r *Robot) Reset() {
	*r = Robot{}
}
