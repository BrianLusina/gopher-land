package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrInvalidNumber       = errors.New("invalid phone number")
	ErrInvalidCode         = errors.New("invalid code")
	ErrInvalidAreaCode     = errors.New("invalid area code")
	ErrInvalidExchangeCode = errors.New("invalid exchange code")
)

// cleanNumber removes all non-digit characters from the phone number
func cleanNumber(number string) string {
	pattern := `[^0-9]`
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(number, "")
}

// exchangeCode returns the exchange code for the given phone number
func exchangeCode(phoneNumber string) (string, error) {
	exchangeCode := phoneNumber[3:6]
	return validNDigit(exchangeCode)
}

func getAreaCode(phoneNumber string) (string, error) {
	exchangeCode := phoneNumber[:3]
	return validNDigit(exchangeCode)
}

func validNDigit(digits string) (string, error) {
	digit := digits[0]
	if digit >= '2' {
		return digits, nil
	}
	return "", ErrInvalidCode
}

func validateNumber(cleanedNumber string) (string, error) {
	if _, e := exchangeCode(cleanedNumber); e != nil {
		return "", ErrInvalidNumber
	}
	if _, e := getAreaCode(cleanedNumber); e != nil {
		return "", ErrInvalidNumber
	}
	return cleanedNumber, nil
}

func Number(phoneNumber string) (string, error) {
	cleanedNumber := cleanNumber(phoneNumber)
	switch {
	case len(cleanedNumber) == 11 && cleanedNumber[0] == '1':
		number := cleanedNumber[1:]
		return validateNumber(number)
	case len(cleanedNumber) == 10:
		return validateNumber(cleanedNumber)
	default:
		return "", ErrInvalidNumber
	}
}

func AreaCode(phoneNumber string) (string, error) {
	cleanedNumber, err := Number(phoneNumber)
	if err != nil {
		return "", ErrInvalidAreaCode
	}
	areaCode := cleanedNumber[:3]
	return validNDigit(areaCode)
}

func Format(phoneNumber string) (string, error) {
	if validNumber, err := Number(phoneNumber); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("(%s) %s-%s", validNumber[:3], validNumber[3:6], validNumber[6:]), nil
	}
}
