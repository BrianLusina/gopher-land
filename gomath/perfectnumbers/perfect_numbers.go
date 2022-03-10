package perfect

import "errors"

var ErrOnlyPositive = errors.New("only positive numbers are allowed")

type Classification string

const (
	ClassificationDeficient Classification = "Deficient"
	ClassificationAbundant  Classification = "Abundant"
	ClassificationPerfect   Classification = "Perfect"
)

// aliquotSum gets the aliquot sum of a number
func aliquotSum(n int64) int64 {
	sum := int64(0)
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	sum := aliquotSum(n)
	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	case sum < n:
		return ClassificationDeficient, nil
	default:
		return "", ErrOnlyPositive
	}
}
