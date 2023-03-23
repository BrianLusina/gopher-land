package car

func main() {
	sportsCarBuilder := NewBuilder("sports")
	familyCarBuilder := NewBuilder("family")

	sportsCar := sportsCarBuilder.Wheels(SportsWheels).TopSpeed(150 * MPH).Build()
	sportsCar.Drive()

	familyCar := familyCarBuilder.Wheels(SteelWheels).TopSpeed(50 * MPH).Build()
	familyCar.Drive()
}
