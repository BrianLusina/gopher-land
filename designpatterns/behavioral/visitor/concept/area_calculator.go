package concept

import (
	"fmt"
	"math"
)

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	area := s.side * s.side
	fmt.Printf("Calculating area for square, %d", area)
	a.area = area
}

func (a *areaCalculator) visitForCircle(c *circle) {
	area := math.Pi * math.Pow(float64(c.radius), 2)
	fmt.Printf("Calculating area for circle, %f", area)
	a.area = int(area)
}

func (a *areaCalculator) visitForRectangle(r *rectangle) {
	area := r.height * r.width
	fmt.Printf("Calculating area for rectangle, %d", area)
	a.area = area
}
