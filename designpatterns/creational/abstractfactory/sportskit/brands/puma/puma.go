package puma

import (
	"gopherland/designpatterns/creational/abstractfactory/sportskit/kit"
)

// Puma is a concrete factory for a Puma sports kit
type PumaKit struct{}

func (a *PumaKit) MakeShoe() kit.ShoeBrand {
	shoe := kit.NewShoe("puma", 14)
	return &PumaShoe{
		Shoe: *shoe,
	}
}

func (a *PumaKit) MakeShirt() kit.ShirtBrand {
	shirt := kit.NewShirt("puma", 14)

	return &PumaShirt{
		Shirt: *shirt,
	}
}
