package deck

import (
	"errors"
	"fmt"
)

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
// Return:
// - (Card) a new card
// - (Error) an error for why it could not create a new card
func NewCard(value int, cardType string) (Card, error) {
	switch cardType {
	case "spades", "hearts", "diamonds", "clubs":
		if value >= 2 && value <= 14 {
			return Card{value, cardType}, nil
		}

		stringValue := fmt.Sprintf("%d", value)
		return Card{}, errors.New(LogMessage("error.card.invalid.value", stringValue))
	default:
		return Card{}, errors.New(LogMessage("error.card.invalid.type", cardType))
	}
}

// NewDeck ...
// This method returns a new deck of cards
// Returns: ([]Card) a slice containing all 52 possible cards in order
func NewDeck() ([]Card, error) {
	cardTypes := []string{"spades", "hearts", "diamonds", "clubs"}

	deckOfCards := []Card{}
	for _, cardType := range cardTypes {
		for value := 2; value <= 14; value++ {
			newCard, err := NewCard(value, cardType)

			if err != nil {
				return nil, err
			}

			deckOfCards = append(deckOfCards, newCard)
		}
	}

	return deckOfCards, nil
}
