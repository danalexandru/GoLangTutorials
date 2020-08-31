package deck

import (
	"fmt"
	"strings"
)

// LogMessage ...
// A map containing all the possible messages returned through the application
var log = map[string]string{
	"error.card.invalid.type":            "Invalid card type \"{0}\". It should be \"spades\", \"hearts\", \"diamonds\" or \"clubs\".",
	"error.card.invalid.value":           "Invalid card value {0}. It should be between 2 and 14.",
	"info.deck.new.deck":                 "New Deck: {0}",
	"info.deck.shuffled.deck":            "Shuffled Deck: {0}",
	"error.deck.shuffle.elements.number": "Inconsistent number of elements after shuffle. Old deck: {0} cards, New deck: {1} cards.",
	"error.deck.shuffle.elements.match":  "The deck was not shuffled. The elements are in the same position",
}

// LogMessage ...
// This function returns the log message with the inner variables replaced
func LogMessage(key string, variables ...string) string {
	message := log[key]
	for i, variable := range variables {
		replace := fmt.Sprintf("{%d}", i)
		message = strings.ReplaceAll(message, replace, variable)
	}

	return message
}
