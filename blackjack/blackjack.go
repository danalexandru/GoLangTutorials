package blackjack

import (
	"fmt"

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
func DrawCard(deckOfCards *[]deck.Card) deck.Card {
	card := (*deckOfCards)[len(*deckOfCards)-1]
	*deckOfCards = (*deckOfCards)[:len(*deckOfCards)-1]

	return card
}

// InitGame ...
// This method creates an "numberOfPlayers" + 1 number of hands (the players + the dealer) and gives them 2 cards from the shuffled deck
// Parameters:
// - numberOfDecks: (int) the number of decks containing distinct 52 cards
// - numberOfPlayers: (int) the number of players for the current game
func InitGame(numberOfDecks int, numberOfPlayers int) ([]deck.Card, Hand, []Hand, error) {

	deckOfCards, err := GenerateBlackjackDeck(numberOfDecks)
	var dealer Hand
	var players []Hand
	if err != nil {
		return nil, nil, nil, err
	}

	// Add the players with empty hands
	for player := 0; player < numberOfPlayers; player++ {
		players = append(players, Hand{})
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < numberOfPlayers; j++ {
			players[j].Push(DrawCard(&deckOfCards))
		}
		dealer.Push(DrawCard(&deckOfCards))
	}

	fmt.Println(players)
	return deckOfCards, dealer, players, nil
}
