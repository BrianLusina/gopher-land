package guns

import "fmt"

func GetGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return NewAk47(), nil
	case "musket":
		return NewMusket(), nil
	default:
		return nil, fmt.Errorf("wrong gun type passed")
	}
}
