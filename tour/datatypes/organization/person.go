package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Type Alias copies fields and method sets and become the exact type
// type TwitterHandler = string

// Type Declaration copies the fields of a type over to a new type
type TwitterHandler string

type Identifiable interface {
	ID() string
}

// Citizen embeds the Identifiable interfaces
type Citizen interface {
	Identifiable
	Country() string
}

type Conflict interface {
	ID() string
}

type Name struct {
	first string
	last  string
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Citizen
}
type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

// NewEuropeanUnionIdentifier creates a new EU identifier, NOTE, that using id with an arbitrary interface offers a lot
// of flexibility, but is not encouraged because it means any type can be used
func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch identifier := id.(type) {
	case string:
		return europeanUnionIdentifier{
			id:      identifier,
			country: country,
		}
	case int:
		return europeanUnionIdentifier{
			id:      strconv.Itoa(identifier),
			country: country,
		}
	case europeanUnionIdentifier:
		return identifier
	default:
		panic("using an invalid type to initialize EU identifier")
	}
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}
func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return eui.country
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
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
	return fmt.Sprintf(p.Citizen.ID())
}

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}
