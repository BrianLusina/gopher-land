package temperature

type (
	// Celsius is a temperature in celsius
	Celsius    float32
	Fahrenheit float32
)

func ToFahrenheit(t Celsius) Fahrenheit {
	return Fahrenheit(t*9/5 + 32)
}
