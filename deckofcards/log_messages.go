package deck

// LogMessage ...
// A map containing all the possible messages returned through the application
var LogMessage = map[string]string{
	"error.card.invalid.type":  "Invalid card type. It should be \"spades\", \"hearts\", \"diamonds\" or \"clubs\".",
	"error.card.invalid.value": "Invalid card value. It should be between 2 and 14.",
}
