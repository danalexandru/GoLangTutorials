package main

import (
	"fmt"
	"strings"

	deck "../deckofcards"
)

// Hand ...
// The current cards that a player / dealer has
type Hand []deck.Card

func (hand Hand) String() string {
	result := make([]string, len(hand))
	for i, card := range hand {
		result[i] = card.String()
	}

	return strings.Join(result, ", ")
}

// CustomString ...
// This is a custom method that will be used by the dealer instead of using the regular "String" method at the beginning
func (hand Hand) CustomString() string {
	return hand[0].String() + ", **HIDDEN**"
}

// Push ...
// This method adds a new card at the end of the current hand
// Parameters:
// - card: (deck.Card) the new card that is added to the current hand
func (hand *Hand) Push(card deck.Card) {
	*hand = append(*hand, card)
}

// Pull ...
// This method takes pulls the last card from the hand
func (hand *Hand) Pull() {
	*hand = (*hand)[:(*hand).Size()-1]
}

// Clear ...
// This method clears the hand completely
func (hand *Hand) Clear() {
	*hand = nil
}

// Size ...
// This method returns the number of cards the current hand has
func (hand Hand) Size() int {
	return len(hand)
}

// ExecTurn ....
// This method is used by a player in order to determine whether or not he wants to hit or stand
// Parameters:
// - deckOfCards: ([]deck.Card) the current deck of cards
func (hand Hand) ExecTurn(deckOfCards []deck.Card) {
	var input string
	for {
		fmt.Printf("What will you do? (h)it or (s)tand: ")
		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			hand.Push(DrawCard(&deckOfCards))
		case "s":
			return
		default:
			fmt.Printf("Invalid string \"%s\". It should be either \"h\" or \"s\".\n", input)
		}
	}
}
