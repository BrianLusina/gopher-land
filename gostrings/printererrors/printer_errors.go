package printererrors

import "fmt"

func PrinterError(s string) string {
	errors := 0
	for _, c := range s {
		if c < 'a' || c > 'm' {
			errors++
		}
	}
	return fmt.Sprintf("%d/%d", errors, len(s))
}
