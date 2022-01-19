package blackjack

import "strings"

type Deck struct {
	Cards map[string]Card
}

type Card struct {
	Value int
}

// MakeDeck returns a deck of cards containing their name(key) and value.
func MakeDeck() Deck {
	deck := Deck{}
	deck.Cards = make(map[string]Card)
	types := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for idx, card := range types {
		switch card {
		case "Jack", "Queen", "King":
			idx = 10
		case "Ace":
			idx = 11
		default:
			idx += 2
		}
		deck.Cards[card] = Card{Value: idx}
	}
	return deck
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	deck := MakeDeck()

	return deck.Cards[strings.Title(strings.ToLower(card))].Value
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	switch {
	case isBlackjack:
		if dealerScore < 10 {
			return "W"
		}
		return "S"
	default:
		return "P"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore >= 17:
		return "S"
	case handScore <= 11:
		return "H"
	case handScore >= 12 && handScore <= 16:
		if dealerScore >= 7 {
			return "H"
		}
		return "S"
	default:
		return "P"
	}
}
