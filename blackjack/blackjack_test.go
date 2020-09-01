package blackjack

import (
	"testing"

	deck "../deckofcards"
)

func TestGenerateBlackjackDeck(t *testing.T) {
	genericDeckOfCards, err := deck.NewDeck()

	if err != nil {
		t.Error(err)
	}

	blackjackDeckOfCards, err := GenerateBlackjackDeck(3)

	if err != nil {
		t.Error(err)
	}

	if len(genericDeckOfCards)*3 != len(blackjackDeckOfCards) {
		t.Errorf("Inconsistency. The new deck should have %d cards. The generated one has %d.", len(genericDeckOfCards)*3, len(blackjackDeckOfCards))
	}
}
