package car

type SportsCarBuilder struct {
	wheels Wheels
	speed  Speed
	color  Color
}

func newSportsCarBuilder() *SportsCarBuilder {
	return &SportsCarBuilder{}
}

func (b *SportsCarBuilder) Wheels(wheels Wheels) Builder {
	b.wheels = wheels
	return b
}

func (b *SportsCarBuilder) TopSpeed(speed Speed) Builder {
	b.speed = speed
	return b
}

func (b *SportsCarBuilder) Paint(color Color) Builder {
	b.color = color
	return b
}

func (b *SportsCarBuilder) Build() ICar {
	return Car{
		wheels: b.wheels,
		speed:  b.speed,
		color:  b.color,
	}
}
