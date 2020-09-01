package deck

import (
	"errors"
	"fmt"
	"testing"
)

type input struct {
	Value         int
	Type          string
	ExpectedCard  Card
	ExpectedError error
}

// "spades", "hearts", "diamonds", "clubs"
var inputs = []input{
	{Value: -1, Type: "gigi", ExpectedCard: Card{}, ExpectedError: errors.New(LogMessage("error.card.invalid.type", "gigi"))},
	{Value: 0, Type: "spades", ExpectedCard: Card{}, ExpectedError: errors.New(LogMessage("error.card.invalid.value", "0"))},
	{Value: 1, Type: "hearts", ExpectedCard: Card{}, ExpectedError: errors.New(LogMessage("error.card.invalid.value", "1"))},
	{Value: 2, Type: "diamonds", ExpectedCard: Card{2, "diamonds"}, ExpectedError: nil},
	{Value: 3, Type: "spades", ExpectedCard: Card{3, "spades"}, ExpectedError: nil},
	{Value: 4, Type: "clubs", ExpectedCard: Card{4, "clubs"}, ExpectedError: nil},
	{Value: 5, Type: "gigi", ExpectedCard: Card{}, ExpectedError: errors.New(LogMessage("error.card.invalid.type", "gigi"))},
	{Value: 6, Type: "spades", ExpectedCard: Card{6, "spades"}, ExpectedError: nil},
	{Value: 7, Type: "hearts", ExpectedCard: Card{7, "hearts"}, ExpectedError: nil},
	{Value: 8, Type: "diamonds", ExpectedCard: Card{8, "diamonds"}, ExpectedError: nil},
	{Value: 9, Type: "clubs", ExpectedCard: Card{9, "clubs"}, ExpectedError: nil},
	{Value: 10, Type: "spades", ExpectedCard: Card{10, "spades"}, ExpectedError: nil},
	{Value: 11, Type: "hearts", ExpectedCard: Card{11, "hearts"}, ExpectedError: nil},
	{Value: 12, Type: "diamonds", ExpectedCard: Card{12, "diamonds"}, ExpectedError: nil},
	{Value: 13, Type: "clubs", ExpectedCard: Card{13, "clubs"}, ExpectedError: nil},
	{Value: 14, Type: "spades", ExpectedCard: Card{14, "spades"}, ExpectedError: nil},
	{Value: 15, Type: "hearts", ExpectedCard: Card{}, ExpectedError: errors.New(LogMessage("error.card.invalid.value", "15"))},
	{Value: 16, Type: "gigi", ExpectedCard: Card{}, ExpectedError: errors.New(LogMessage("error.card.invalid.type", "gigi"))},
}

// TestNewCard ...
// This method tests the functionality of the "NewCard" method from "card.go"
func TestNewCard(t *testing.T) {
	for testIndex, item := range inputs {
		card, err := NewCard(item.Value, item.Type)

		if (card != item.ExpectedCard) ||
			(err != nil && item.ExpectedError != nil && err.Error() != item.ExpectedError.Error()) {
			t.Errorf("Invalid test: %d) {%v}\n"+
				"\t- Card: {expected: %v,\t received: %v}\n"+
				"\t- Error: {expected: %v,\t received: %v}",
				testIndex, item, item.ExpectedCard, card, item.ExpectedError, err)
		}

	}
}

// TestNewDeck ...
// This method tests the functionality of the "NewDeck" method from "card.go"
func TestNewDeck(t *testing.T) {
	_, err := NewDeck()

	if err != nil {
		t.Error(err)
	}
}

// TestShuffleDeck ...
// This method tests the functionality of the "ShuffleDeck" method from "card.go"
func TestShuffleDeck(t *testing.T) {
	deckOfCards, err := NewDeck()
	shuffledDeckOfCards, _ := NewDeck()

	if err != nil {
		t.Error(err)
	}

	copy(shuffledDeckOfCards[:], deckOfCards)

	ShuffleDeck(shuffledDeckOfCards)

	if len(deckOfCards) != len(shuffledDeckOfCards) {
		t.Error(LogMessage("error.deck.shuffle.elements.number",
			fmt.Sprintf("%d", len(deckOfCards)),
			fmt.Sprintf("%d", len(shuffledDeckOfCards))))
	}

	isEqual := func(deck1, deck2 []Card) bool {
		for i := range deck1 {
			if deck1[i] != deck2[i] {
				return false
			}
		}

		return true
	}

	numberOfAttempts := 1
	for isEqual(deckOfCards, shuffledDeckOfCards) {
		if numberOfAttempts == 3 {
			t.Error(LogMessage("error.deck.shuffle.elements.match"))
			break
		} else {
			ShuffleDeck(shuffledDeckOfCards)
			numberOfAttempts++
		}
	}
}

// TestSortDeck ...
// This method tests the functionality of the "SortDeck" method from "card.go"
func TestSortDeck(t *testing.T) {
	deckOfCards, err := NewDeck()
	sortedDeckOfCards, _ := NewDeck()

	if err != nil {
		t.Error(err)
	}

	copy(sortedDeckOfCards[:], deckOfCards)

	ShuffleDeck(deckOfCards)
	SortDeck(sortedDeckOfCards)

	if len(deckOfCards) != len(sortedDeckOfCards) {
		t.Error(LogMessage("error.deck.sort.elements.number",
			fmt.Sprintf("%d", len(deckOfCards)),
			fmt.Sprintf("%d", len(sortedDeckOfCards))))
	}

	cardTypes := []string{"spades", "hearts", "clubs", "diamonds"}
	getIndexOfCardType := func(cardType string) int {
		for index, value := range cardTypes {
			if value == cardType {
				return index
			}
		}

		return -1
	}

	for i := 0; i < len(sortedDeckOfCards)-1; i++ {
		switch {
		case getIndexOfCardType(sortedDeckOfCards[i].Type) > getIndexOfCardType(sortedDeckOfCards[i+1].Type):
			t.Error(LogMessage("error.deck.sort.elements.by.type", sortedDeckOfCards[i].Type, sortedDeckOfCards[i+1].Type))
		case getIndexOfCardType(sortedDeckOfCards[i].Type) == getIndexOfCardType(sortedDeckOfCards[i+1].Type) &&
			sortedDeckOfCards[i].Value > sortedDeckOfCards[i+1].Value:
			t.Error(LogMessage("error.deck.sort.elements.by.type", sortedDeckOfCards[i].Type, sortedDeckOfCards[i+1].Type))
		}
	}

}

func TestString(t *testing.T) {
	// "spades", "hearts", "diamonds", "clubs"
	type stringInput struct {
		card           Card
		expectedString string
	}
	testCards := []stringInput{
		{card: Card{Value: 2, Type: "spades"}, expectedString: "2 of Spades"},
		{card: Card{Value: 4, Type: "hearts"}, expectedString: "4 of Hearts"},
		{card: Card{Value: 6, Type: "diamonds"}, expectedString: "6 of Diamonds"},
		{card: Card{Value: 8, Type: "clubs"}, expectedString: "8 of Clubs"},
		{card: Card{Value: 10, Type: "spades"}, expectedString: "10 of Spades"},
		{card: Card{Value: 11, Type: "hearts"}, expectedString: "Ace of Hearts"},
		{card: Card{Value: 12, Type: "diamonds"}, expectedString: "Jack of Diamonds"},
		{card: Card{Value: 13, Type: "clubs"}, expectedString: "Queen of Clubs"},
		{card: Card{Value: 14, Type: "spades"}, expectedString: "King of Spades"},
	}

	for index, test := range testCards {
		if test.card.String() != test.expectedString {
			t.Errorf("Test %d failed. The result should have been %s. Instead we got %s.", index, test.expectedString, test.card.String())
		}
	}
}
