// Package weather implements a library for forecasting weather condition for a city.
package weather

// CurrentCondition represents the current weather condition for a city.
var CurrentCondition string

// CurrentLocation is the current geographic location
var CurrentLocation string

// Forecast returns the weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
