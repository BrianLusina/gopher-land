package guns

type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47",
			power: 4,
		},
	}
}
