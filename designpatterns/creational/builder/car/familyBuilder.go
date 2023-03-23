package car

type FamilyCarBuilder struct {
	wheels Wheels
	speed  Speed
	color  Color
}

func newFamilyCarBuilder() *SportsCarBuilder {
	return &SportsCarBuilder{}
}

func (b *FamilyCarBuilder) Wheels(wheels Wheels) Builder {
	b.wheels = wheels
	return b
}

func (b *FamilyCarBuilder) TopSpeed(speed Speed) Builder {
	b.speed = speed
	return b
}

func (b *FamilyCarBuilder) Paint(color Color) Builder {
	b.color = color
	return b
}

func (b *FamilyCarBuilder) Build() ICar {
	return Car{
		wheels: b.wheels,
		speed:  b.speed,
		color:  b.color,
	}
}
