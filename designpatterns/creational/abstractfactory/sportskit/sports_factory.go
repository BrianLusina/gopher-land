package sportskit

import (
	"fmt"
	"gopherland/designpatterns/creational/abstractfactory/sportskit/brands/adidas"
	"gopherland/designpatterns/creational/abstractfactory/sportskit/brands/nike"
	"gopherland/designpatterns/creational/abstractfactory/sportskit/brands/puma"
	"gopherland/designpatterns/creational/abstractfactory/sportskit/kit"
)

// SportsFactory is an abstract factory interface
type SportsFactory interface {
	MakeShoe() kit.ShoeBrand
	MakeShirt() kit.ShirtBrand
}

func GetSportsFactory(brand string) (SportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidas.AdidasKit{}, nil
	case "nike":
		return &nike.NikeKit{}, nil
	case "puma":
		return &puma.PumaKit{}, nil
	default:
		return nil, fmt.Errorf("unknown brand")
	}
}
