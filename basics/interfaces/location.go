package interfaces

import (
	"context"
	"errors"
)

type locator interface {
	getCoordinates(ctx context.Context, address string) (lat, lng float32, err error)
}

type locate struct {
}

func (l locate) getCoordinates(ctx context.Context, address string) (lat, lng float32, err error) {
	isValid := l.validateAddress(address)
	if !isValid {
		return 0, 0, errors.New("invalid address")
	}

	if err := ctx.Err(); err != nil {
		return 0, 0, err
	}

	// get address location
	return 1, 1, nil
}

func (l locate) validateAddress(address string) bool {
	return address != ""
}
