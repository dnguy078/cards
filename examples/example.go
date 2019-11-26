package main

import (
	"fmt"

	"github.com/dnguy078/cards"
)

func main() {
	d := cards.NewDeck()
	d.Shuffle()
	c, _ := d.DealOneCard()
	fmt.Println(c)
}
