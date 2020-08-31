package deck

import (
	"fmt"
	"strings"
)

// LogMessage ...
// A map containing all the possible messages returned through the application
var log = map[string]string{
	"error.card.invalid.type":  "Invalid card type \"{0}\". It should be \"spades\", \"hearts\", \"diamonds\" or \"clubs\".",
	"error.card.invalid.value": "Invalid card value {0}. It should be between 2 and 14.",
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
