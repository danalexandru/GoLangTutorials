package blackjack

import (
	deck "../deckofcards"
)

// GenerateBlackjackDeck ...
// This method returns a number of decks concatenated together to be used in the blackjack game
// Parameters:
// - numberOfDecks: (int) the number of decks containing distinct 52 cards
// Return:
// - ([]deck.Card) a bigger deck containing "numberOfDecks" decks concatenated together
// - (Error) if an error occurred
func GenerateBlackjackDeck(numberOfDecks int) ([]deck.Card, error) {
	deckOfCards := []deck.Card{}

	for i := 0; i < numberOfDecks; i++ {
		tempDeck, err := deck.NewDeck()

		if err != nil {
			return nil, err
		}

		deckOfCards = append(deckOfCards, tempDeck...)
	}

	deck.ShuffleDeck(deckOfCards)

	return deckOfCards, nil
}

// DrawCard ...
// This method extracts the last card from the current deck of cards
// Parameters:
// - deckOfCards: ([]deck.Card) the current deck of cards
// Return:
// - (deck.Card) the last card from the deckOfCards (now the deckOfCards will not have that card anymore, since it was drawn out)
func DrawCard(deckOfCards []deck.Card) deck.Card {
	card := deckOfCards[len(deckOfCards)-1]
	deckOfCards = deckOfCards[:len(deckOfCards)-1]

	return card
}
