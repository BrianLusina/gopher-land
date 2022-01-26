package lasagna

// PreparationTime returns the total preparation time for the given layers.
// The time is calculated as the average preparation time of each layer.
// If the average preparation time is not defined, the default time 2 is used.
func PreparationTime(layers []string, time int) int {
	l := len(layers)
	if time == 0 {
		time = 2
		return time * l
	}
	return time * l
}

// Quantities returns the quantities of noodles and sauce needed to make the lasanga.
func Quantities(layers []string) (int, float64) {
	n := 50
	s := 0.2
	var noodles int
	var sauce float64

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += n
		}
		if layer == "sauce" {
			sauce += s
		}
	}
	return noodles, sauce
}

// AddSecretIngredient adds a secret ingredient to the lasagna in myList from the friendsList.
func AddSecretIngredient(friendsList, myList []string) {
	last := len(friendsList) - 1
	secret := friendsList[last]
	myList[len(myList)-1] = secret
}

// ScaleRecipe scales the recipe by the given number of portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	if portions == 0 {
		return quantities
	}

	amounts := make([]float64, len(quantities))
	multiplier := float64(portions) / float64(2)

	for idx, qty := range quantities {
		qty *= multiplier
		amounts[idx] = qty
	}

	return amounts
}
