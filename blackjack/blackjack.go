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
func GetWinner(players []Hand, dealer Hand) Hand {
	var bestPlayer Hand
	bestPlayerIndex := 0
	bestPlayer = players[0]

	for i := 1; i < len(players); i++ {
		if bestPlayer.GetBiggestScore() < players[i].GetBiggestScore() {
			bestPlayer = players[i]
			bestPlayerIndex = i
		}
	}

	bestPlayerScore := bestPlayer.GetBiggestScore()
	dealerScore := dealer.GetBiggestScore()

	fmt.Printf("\nWinner:\n")
	switch {
	case bestPlayerScore > dealerScore:
		fmt.Printf("\tPlayer no #%d won. Final score: %d.\n", (bestPlayerIndex + 1), bestPlayer.GetBiggestScore())
		return bestPlayer
	case bestPlayerScore < dealerScore:
		fmt.Printf("\tDealer won. Final score: %d.\n", dealerScore)
		return dealer
	case bestPlayerScore == dealerScore && bestPlayerScore == -1:
		fmt.Printf("\tAll players and the dealer have BUSTED. (:.\n")
		return Hand{}
	default:
		fmt.Printf("\tTie between player no #%d and the dealer. Final Score: %d.\n", (bestPlayerIndex + 1), bestPlayerScore)
		return bestPlayer
	}
}

func main() {
	deckOfCards, dealer, players, err := InitGame(3, 3)

	if err != nil {
		fmt.Println(err)
		return
	}

	for len(deckOfCards) != 0 {
		fmt.Println("---------------------------")
		fmt.Println("Current game status: ")
		fmt.Printf("Dealer: { %s } \n", dealer.CustomString())
		for i, player := range players {
			fmt.Printf("Player no #%d: { { %s }, { Score: %s } }\n", (i + 1), player.String(), player.FormatedScore())
		}
		fmt.Println("---------------------------")

		for i, player := range players {
			fmt.Println("-----")
			fmt.Printf("Player no #%d)\n", (i + 1))
			player.ExecTurn(&deckOfCards)
			fmt.Println("-----")
			fmt.Println()
		}
		dealer.ExecDealerTurn(&deckOfCards)

		GetWinner(players, dealer)
	}

	// players[0].ExecTurn(&deckOfCards)
	// dealer.ExecDealerTurn(&deckOfCards)

	fmt.Println("Dealer Hand: ", dealer.String())
	if dealer.GetBiggestScore() != -1 {
		fmt.Println("Dealer Score: ", dealer.FormatedScore())
	} else {
		fmt.Println("Dealer Busted")
	}
}
