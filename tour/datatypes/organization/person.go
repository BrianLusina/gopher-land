package organization

import (
	"errors"
	"fmt"
	"strings"
)

// Type Alias copies fields and method sets and become the exact type
// type TwitterHandler = string

// Type Declaration copies the fields of a type over to a new type
type TwitterHandler string

type Identifiable interface {
	ID() string
}

type Name struct {
	first string
	last  string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Identifiable {
	return socialSecurityNumber(
		value
	)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Identifiable
}

func NewPerson(firstName, lastName string, identifiable Identifiable) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Identifiable: identifiable,
	}
}

// Reference based over value type with pointer (*) to allow referencing the same object instead of using a copy
func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}
	p.twitterHandler = handler
	return nil
}

func (p *Person) ID() string {
	return fmt.Sprintf(p.Identifiable.ID())
}

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}
