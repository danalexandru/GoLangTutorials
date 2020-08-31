package deck

import "errors"

// Card ...
// Structure used for representing one card from a deck
type Card struct {
	Value int
	Type  string
}

// NewCard ...
// This method returns a new card, or an error should the input values be invalid
// Parameters:
// - value: (int) the value of the card (between 2 and 14)
// - cardType: (string) the type of the card (valid values: {"spades", "hearts", "diamonds", "clubs"})
func NewCard(value int, cardType string) (Card, error) {
	switch cardType {
	case "spades", "hearts", "diamonds", "clubs":
		if value >= 2 && value <= 14 {
			return Card{value, cardType}, nil
		}
		return Card{}, errors.New(LogMessage["error.card.invalid.value"])
	default:
		return Card{}, errors.New(LogMessage["error.card.invalid.type"])
	}
}
