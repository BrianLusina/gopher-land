// Package leap implements a simple library to get leap years
package leap

// IsLeapYear determines if a year is leap
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
