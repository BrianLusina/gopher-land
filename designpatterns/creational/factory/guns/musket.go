package guns

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket",
			power: 1,
		},
	}
}
