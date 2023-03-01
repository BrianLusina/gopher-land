package nike

import (
	"gopherland/designpatterns/creational/abstractfactory/sportskit/kit"
)

// Nike is a concrete factory for a Nike sports kit
type NikeKit struct{}

func (a *NikeKit) MakeShoe() kit.ShoeBrand {
	shoe := kit.NewShoe("nike", 14)
	return &NikeShoe{
		Shoe: *shoe,
	}
}

func (a *NikeKit) MakeShirt() kit.ShirtBrand {
	shirt := kit.NewShirt("nike", 14)

	return &NikeShirt{
		Shirt: *shirt,
	}
}
