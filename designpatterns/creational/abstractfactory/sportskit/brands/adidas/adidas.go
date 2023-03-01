package adidas

import "gopherland/designpatterns/creational/abstractfactory/sportskit/kit"

// AdidasKit is a concrete factory for an Adidas sports kit
type AdidasKit struct{}

func (a *AdidasKit) MakeShoe() kit.ShoeBrand {
	shoe := kit.NewShoe("adidas", 14)

	return &AdidasShoe{
		Shoe: *shoe,
	}
}

func (a *AdidasKit) MakeShirt() kit.ShirtBrand {
	shirt := kit.NewShirt("adidas", 14)

	return &AdidasShirt{
		Shirt: *shirt,
	}
}
