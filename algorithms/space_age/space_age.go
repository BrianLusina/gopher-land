package space

type Planet string

const (
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Earth   Planet = "Earth"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

// EarthOrbitalSeconds is the number of seconds it takes the Earth to orbit the Sun
const EarthOrbitalSeconds = 31557600

var planetPeriods = map[Planet]float64{
	Mercury: 0.2408467 * EarthOrbitalSeconds,
	Venus:   0.61519726 * EarthOrbitalSeconds,
	Earth:   1.0 * EarthOrbitalSeconds,
	Mars:    1.8808158 * EarthOrbitalSeconds,
	Jupiter: 11.862615 * EarthOrbitalSeconds,
	Saturn:  29.447498 * EarthOrbitalSeconds,
	Uranus:  84.016846 * EarthOrbitalSeconds,
	Neptune: 164.79132 * EarthOrbitalSeconds,
}

// Age takes in seconds and a planet type and calculates the number of years
func Age(seconds float64, planet Planet) float64 {
	earthOrbitalPeriod, ok := planetPeriods[planet]

	if !ok {
		earthOrbitalPeriod = 1
	}

	return seconds / earthOrbitalPeriod
}
