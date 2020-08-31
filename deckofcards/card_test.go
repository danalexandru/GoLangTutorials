package deck

import (
	"errors"
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
