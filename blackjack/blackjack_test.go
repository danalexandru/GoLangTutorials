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

func TestDrawCard(t *testing.T) {
	deckOfCards, err := GenerateBlackjackDeck(3)

	if err != nil {
		t.Error(err)
	}

	numberOfCardsInDeck := len(deckOfCards)
	lastCard := deckOfCards[numberOfCardsInDeck-1]
	drawnCard := DrawCard(deckOfCards)

	if len(deckOfCards) != numberOfCardsInDeck {
		t.Errorf("The deck should have %d after the last card was drawn. Instead, the deck has %d.", numberOfCardsInDeck-1, len(deckOfCards))
	}

	if lastCard != drawnCard {
		t.Errorf("The card drawn should have been %v. Instead, the \"DrawCard\" method got %v.", lastCard, drawnCard)
	}
}
