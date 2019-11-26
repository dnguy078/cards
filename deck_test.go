package cards

import (
	"testing"
)

func TestDeck_Shuffle(t *testing.T) {
	testMap := make(map[*Card]bool)
	d := NewDeck()
	d.Shuffle()

	for c, _ := d.DealOneCard(); c != nil; c, _ = d.DealOneCard() {
		if testMap[c] {
			t.Errorf("duplicate card dealt in stack, card: %s", c)
		} else {
			testMap[c] = true
		}
	}

	if len(testMap) != 52 {
		t.Errorf("Got: %d , Expected: 52 cards to be dealt", len(testMap))
	}
}

func TestDeck_DealOneCard(t *testing.T) {
	d := NewDeck()
	for i := 0; i < 52; i++ {
		_, err := d.DealOneCard()
		if err != nil {
			t.Error("should not return error when dealing filled deck")
		}
	}

	// attempt to deal addictional card
	c, err := d.DealOneCard()
	if c != nil && err == nil {
		t.Error("expected to return nil card and error")
	}
}
