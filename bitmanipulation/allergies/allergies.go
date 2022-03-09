package allergies

var allergyScores = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	var result []string
	for allergen := range allergyScores {
		if AllergicTo(allergies, allergen) {
			result = append(result, allergen)
		}
	}
	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	return allergyScores[allergen]&allergies > 0
}
