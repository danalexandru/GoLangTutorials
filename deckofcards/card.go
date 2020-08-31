package deck

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
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
	cardTypes := []string{"spades", "hearts", "clubs", "diamonds"}

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

	// fmt.Println(LogMessage("info.deck.new.deck", fmt.Sprintf("%v", deckOfCards)))
	return deckOfCards, nil
}

// ShuffleDeck ...
// This method shuffles the card of the current deck
// Parameters:
// - deckOfCards: ([]Card) The current deck of cards, presumably unshuffled
func ShuffleDeck(deckOfCards []Card) {
	rand.Shuffle(len(deckOfCards), func(i, j int) { deckOfCards[i], deckOfCards[j] = deckOfCards[j], deckOfCards[i] })
	// fmt.Println(LogMessage("info.deck.shuffled.deck", fmt.Sprintf("%v", deckOfCards)))
}

// SortDeck ...
// This method sorts the cards of a deck (poresumably shuffled) in ascending order
// Parameters:
// - deckOfCards (sort.Interface) The deck of cards that needs to be sorted
func SortDeck(deckOfCards []Card) {
	cardTypes := []string{"spades", "hearts", "clubs", "diamonds"}
	getIndexOfCardType := func(cardType string) int {
		for index, value := range cardTypes {
			if value == cardType {
				return index
			}
		}

		return -1
	}

	sort.SliceStable(deckOfCards, func(i, j int) bool {
		switch {
		case getIndexOfCardType(deckOfCards[i].Type) == -1 || getIndexOfCardType(deckOfCards[j].Type) == -1:
			return false
		case getIndexOfCardType(deckOfCards[i].Type) < getIndexOfCardType(deckOfCards[j].Type):
			return true
		case getIndexOfCardType(deckOfCards[i].Type) == getIndexOfCardType(deckOfCards[j].Type) &&
			deckOfCards[i].Value < deckOfCards[j].Value:
			return true
		default:
			return false
		}
	})
}
