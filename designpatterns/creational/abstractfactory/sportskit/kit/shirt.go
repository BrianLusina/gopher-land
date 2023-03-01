package kit

// ShirtBrand is an Abstract product representing a shirt
type ShirtBrand interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Shirt struct {
	logo string
	size int
}

// NewShirt creates a new Shoe
func NewShirt(logo string, size int) *Shirt {
	return &Shirt{
		logo: logo,
		size: size,
	}
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) GetLogo() string {
	return s.logo
}

func (s *Shirt) SetSize(size int) {
	s.size = size
}

func (s *Shirt) GetSize() int {
	return s.size
}
