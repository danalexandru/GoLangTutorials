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

// Score ...
// This method returns the score of the current player
func (hand Hand) Score() []int {
	result := []int{0}
	// calculate all possible results
	for _, card := range hand {
		if card.Value == 11 {
			result = append(result, result...)
			for i := 0; i < len(result)/2; i++ {
				result[i]++
				result[len(result)/2+i] += 11
			}
		} else if card.Value >= 12 {
			for i := 0; i < len(result); i++ {
				result[i] += 10
			}
		} else {
			for i := 0; i < len(result); i++ {
				result[i] += card.Value
			}
		}
	}

	// filter out scores > 21
	for i := 0; i < len(result)-1; {
		if result[i] > 21 {
			result = append(result[:i], result[i+1:]...)
		} else {
			i++
		}
	}
	if result[len(result)-1] > 21 {
		result = result[:len(result)-1]
	}

	return result
}

// FormatedScore ...
// This method formats the score in a user-firendly manner
// Parameters:
// - score: ([]int) the soft scores returned by the "Score" method
func (hand Hand) FormatedScore() string {
	score := hand.Score()

	switch len(score) {
	case 0:
		return fmt.Sprintf("BUST")
	case 1:
		if score[0] == 21 {
			return "BLACKJACK"
		}
		return fmt.Sprintf("%d", score[0])

	default:
		result := fmt.Sprintf("%d", score[0])
		for i := 1; i < len(score); i++ {
			if score[i] == 21 {
				result = fmt.Sprintf("%s, BLACKJACK", result)
			} else {
				result = fmt.Sprintf("%s, %d", result, score[i])
			}
		}
		return "[ " + result + " ]"
	}

}

// GetBiggestScore ...
// This method returns the greathest score of the current player
func (hand Hand) GetBiggestScore() int {
	allScores := hand.Score()

	switch len(allScores) {
	case 0:
		return -1 // Bust
	default:
		maxScore := -1
		for _, softScore := range allScores {
			if softScore > maxScore {
				maxScore = softScore
			}
		}

		return maxScore
	}
}

// ExecTurn ....
// This method is used by a player in order to determine whether or not he wants to hit or stand
// Parameters:
// - deckOfCards: ([]deck.Card) the current deck of cards
func (hand *Hand) ExecTurn(deckOfCards []deck.Card) {
	var input string
	for {
		fmt.Println()
		fmt.Printf("\t- Current cards: { %s }\n", hand.String())
		fmt.Printf("\t- Current score: %s\n", hand.FormatedScore())
		fmt.Println()

		fmt.Printf("What will you do? (h)it or (s)tand: ")
		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			hand.Push(DrawCard(&deckOfCards))
		case "s":
			fmt.Println()
			return
		default:
			fmt.Printf("Invalid string \"%s\". It should be either \"h\" or \"s\".\n", input)
		}

		if len(hand.Score()) == 0 {
			fmt.Printf("BUST! Your current cards are: { %s }\n", hand.String())
			return
		}
	}
}

// ExecDealerTurn ...
// This method executes the dealer's turn
// Parameters:
// - deckOfCards: ([]deck.Card) the current deck of cards
func (hand *Hand) ExecDealerTurn(deckOfCards []deck.Card) {
	for hand.GetBiggestScore() <= 16 || (hand.GetBiggestScore() == 17 && len(hand.Score()) > 1) {
		hand.Push(DrawCard(&deckOfCards))

		if hand.GetBiggestScore() == -1 {
			return
		}
	}
}
