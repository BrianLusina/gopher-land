package funcs

import (
	"fmt"
)

func main() {
	result, err := Divide(6, 3)

	if err != nil {
		fmt.Println("Division result", result)
	} else {
		fmt.Println("Error occurred", err.Error())
	}

	numbers := []float64{1, 3}
	fmt.Println(add(numbers...))
}
