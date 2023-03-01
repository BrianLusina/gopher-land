package kit

// ShoeBrand is an Abstract product representing a shoe
type ShoeBrand interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Shoe struct {
	logo string
	size int
}

// NewShoe creates a new Shoe
func NewShoe(logo string, size int) *Shoe {
	return &Shoe{
		logo: logo,
		size: size,
	}
}

func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}

func (s *Shoe) GetSize() int {
	return s.size
}
