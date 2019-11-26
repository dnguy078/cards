package cards

import "fmt"

// Card represents a playing card with type and suit
type Card struct {
	Type string
	Suit string
}

func (c *Card) String() string {
	return fmt.Sprintf("%s_%s", c.Type, c.Suit)
}
