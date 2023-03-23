package car

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type (
	Builder interface {
		Paint(color Color) Builder
		Wheels(wheels Wheels) Builder
		TopSpeed(speed Speed) Builder
		Build() ICar
	}

	ICar interface {
		Drive() error
		Stop() error
	}
)

func NewBuilder(builderType string) Builder {
	switch builderType {
	case "sports":
		return newSportsCarBuilder()
	case "family":
		return newFamilyCarBuilder()
	default:
		return nil
	}
}
