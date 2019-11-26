package cards

import (
	"errors"
	"math/rand"
	"time"
)

// Deck represents deck of cards
type Deck struct {
	cards []*Card
}

// NewDeck returns aDeckd
func NewDeck() *Deck {
	d := make([]*Card, 0)

	types := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	suits := []string{"Heart", "Diamond", "Club", "Spade"}

	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			c := &Card{
				Type: types[i],
				Suit: suits[n],
			}
			d = append(d, c)
		}
	}
	return &Deck{
		cards: d,
	}
}

// Shuffle shuffles the poker deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(d.cards); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d.cards[r], d.cards[i] = d.cards[i], d.cards[r]
		}
	}
}

// DealOneCard returns one card from d.cardsk
func (d *Deck) DealOneCard() (*Card, error) {
	if len(d.cards) == 0 {
		return nil, errors.New("empty deck")
	}
	dealt := d.cards[0]
	d.cards = d.cards[1:]
	return dealt, nil
}
