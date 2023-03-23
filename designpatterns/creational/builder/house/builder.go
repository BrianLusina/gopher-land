package house

type Builder interface {
	setWindowType()
	setDoorType()
	setNumberOfFloors()
	getHouse() House
}

func getBuilder(builderType string) Builder {
	switch builderType {
	case "normal":
		return newNormalBuilder()
	case "igloo":
		return newIglooBuilder()
	default:
		return nil
	}
}
