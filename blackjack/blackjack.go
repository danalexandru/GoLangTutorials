package blackjack

import deck "../deckofcards"

// GenerateBlackjackDeck ...
// This method returns a number of decks concatenated together to be used in the blackjack game
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
