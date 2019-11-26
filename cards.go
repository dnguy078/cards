package cards

import "fmt"

// Card represents a playing card with type and suit
type Card struct {
	Type string
	Suit string
}

// String returns the string representation of a Card
func (c *Card) String() string {
	return fmt.Sprintf("%s_%s", c.Type, c.Suit)
}
