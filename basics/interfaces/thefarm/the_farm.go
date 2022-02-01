package thefarm

import (
	"errors"
	"fmt"
)

// SillyNephewError type here.
type SillyNephewError struct {
	Cows int
}

func (sn *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", sn.Cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	switch {
	case err != nil && err != ErrScaleMalfunction:
		return 0, err
	case err == ErrScaleMalfunction && fodder > 0:
		return (fodder * 2) / float64(cows), nil
	case fodder < 0:
		return 0, errors.New("negative fodder")
	case cows == 0:
		return 0, errors.New("division by zero")
	case cows < 0:
		return 0, &SillyNephewError{Cows: cows}
	default:
		return fodder / float64(cows), nil
	}
}
