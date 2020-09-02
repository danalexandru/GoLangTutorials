package main

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
// Return:
// - ([]deck.Card) the current deck of cards
// - (Hand) the dealer
// - ([]Hand) the players
// - (error) if an error occurred
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

	return deckOfCards, dealer, players, nil
}

// GetWinner ...
// This method returns the winner of the game
// Parameters:
// players: ([]Hand) all the players of the current game
// dealer: (Hand) the dealer of the current game
// Return:
// (Hand) the winner of the game
func GetWinner(players []Hand, dealer Hand) {
	var bestPlayer Hand
	bestPlayerIndex := 0
	copy(bestPlayer, players[0])
	for i := 1; i < len(players); i++ {
		if bestPlayer.GetBiggestScore() > players[i].GetBiggestScore() {
			copy(bestPlayer, players[i])
			bestPlayerIndex = i
		}
	}

	bestPlayerScore := bestPlayer.GetBiggestScore()
	dealerScore := dealer.GetBiggestScore()

	switch {
	case bestPlayerScore > dealerScore:
		fmt.Printf("Player no #%d won. Final score: %d.\n", bestPlayerIndex, bestPlayer.GetBiggestScore())
	case bestPlayerScore < dealerScore:
		fmt.Printf("Dealer won. Final score: %d.\n", dealerScore)
	case bestPlayerScore == dealerScore && bestPlayerScore == -1:
		fmt.Printf("All players and the dealer have BUSTED. (:.\n")
	default:
		fmt.Printf("Tie between player no #%d and the dealer. Final Score: %d.\n", bestPlayerIndex, bestPlayerScore)
	}
}

func main() {
	deckOfCards, dealer, players, err := InitGame(3, 3)

	if err != nil {
		fmt.Println(err)
		return
	}

	players[0].ExecTurn(deckOfCards, players, dealer)

}
