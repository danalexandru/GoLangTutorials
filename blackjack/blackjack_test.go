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

func TestString(t *testing.T) {
	// "spades", "hearts", "diamonds", "clubs"
	hand := Hand{
		deck.Card{Value: 2, Type: "spades"},
		deck.Card{Value: 4, Type: "hearts"},
		deck.Card{Value: 6, Type: "diamonds"},
		deck.Card{Value: 8, Type: "clubs"},
		deck.Card{Value: 10, Type: "spades"},
		deck.Card{Value: 11, Type: "hearts"},
		deck.Card{Value: 12, Type: "diamonds"},
		deck.Card{Value: 13, Type: "clubs"},
		deck.Card{Value: 14, Type: "spades"},
	}

	expectedResult := "2 of Spades, " +
		"4 of Hearts, " +
		"6 of Diamonds, " +
		"8 of Clubs, " +
		"10 of Spades, " +
		"Ace of Hearts, " +
		"Jack of Diamonds, " +
		"Queen of Clubs, " +
		"King of Spades"

	if hand.String() != expectedResult {
		t.Errorf("Inconsistent conversion from Hand to string.\n"+
			"\t- Expected result: \"%s\".\n"+
			"\t- Received result: \"%s\".",
			expectedResult, hand.String())
	}
}

func TestHand(t *testing.T) {
	hand := Hand{}

	hand.Push(deck.Card{Value: 2, Type: "spades"})
	hand.Push(deck.Card{Value: 4, Type: "spades"})
	hand.Push(deck.Card{Value: 6, Type: "spades"})

	if hand.Size() != 3 {
		t.Errorf("Current hand has %d. Elements. It should have %d.", hand.Size(), 3)
	}

	hand.Pull()

	if hand.Size() != 2 {
		t.Errorf("Current hand has %d. Elements. It should have %d.", hand.Size(), 2)
	}

	hand.Clear()

	if hand.Size() != 0 {
		t.Errorf("Hand should be empty. Instead it has %d elements.", hand.Size())
	}
}
