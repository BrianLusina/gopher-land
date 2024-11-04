// Package main contains functions to calculate and map points on a polar coordinate system to a cartesian coordinate system.
package goroutines

import (
	"fmt"
	"math"
	"runtime"
)

type polar struct {
	radius float64
	theta  float64
}

type cartesian struct {
	x float64
	y float64
}

const result = "Polar: radius=%.02f angle=%.02f degrees -- Cartesian: x=%.02f y=%.02f\n\n"

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, or %s to quit:"

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoordinate := <-questions
			theta := polarCoordinate.theta * math.Pi / 180.0
			cartesian := cartesian{
				x: polarCoordinate.radius * math.Cos(theta),
				y: polarCoordinate.radius * math.Sin(theta),
			}
			answers <- cartesian
		}
	}()
	return answers
}

func interact(questions chan polar, answers chan cartesian) {
	for {
		p := polar{}
		// Ask the user for a radius and angle.
		// Note: we're using fmt.Scanf here, but this is just for example purposes.
		// In a real program, you'd use a library like "github.com/mattn/go-scanf".
		fmt.Println("Enter a radius and an angle (in degrees), e.g., 12.5 90, or q to quit:")
		_, err := fmt.Scanf("%f %f", &p.radius, &p.theta)
		if err != nil {
			break
		}
		if p.radius < 0 {
			fmt.Println("The radius must be non-negative.")
			continue
		}
		if p.theta < 0 || p.theta > 360 {
			fmt.Println("The angle must be in the range [0, 360).")
			continue
		}
		// Send the polar coordinates to the solver.
		questions <- p
		// Receive the answer from the solver.
		a := <-answers
		// Display the answer.
		fmt.Printf(result, p.radius, p.theta, a.x, a.y)
	}
	fmt.Println("Bye!")
}

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+C")
	}
}

func run_polar_to_cartesian() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}
