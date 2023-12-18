package validate

import (
	"fmt"

	"github.com/HarshitCd/GoPoker/deck"
)

type CheckHand struct{}

func New() CheckHand {
	return CheckHand{}
}

func (ch *CheckHand) Validate(h []deck.Card, flop []deck.Card) {
	f := append(h, flop...)

	if ch.StraightFlush(f) {
		fmt.Println("Straight Flush")
	} else if ch.FourOfAKind(f) {
		fmt.Println("Four Of A Kind")
	} else if ch.FullHouse(f) {
		fmt.Println("Full House")
	} else if ch.Flush(f) {
		fmt.Println("Flush")
	} else if ch.Straight(f) {
		fmt.Println("Straight")
	} else if ch.ThreeOfAKind(f) {
		fmt.Println("Three Of A Kind")
	} else if ch.TwoPair(f) {
		fmt.Println("Two Pair")
	} else if ch.OnePair(f) {
		fmt.Println("One Pair")
	} else {
		fmt.Println("No Hand")
	}
}

func (ch *CheckHand) StraightFlush(h []deck.Card) bool {
	var suits []deck.Suit = []deck.Suit{deck.Spades, deck.Clubs, deck.Diamonds, deck.Hearts}

	for _, suit := range suits {
		var values [15]uint16
		for _, c := range h {
			if c.Suit == suit {
				values[c.Value] += 1
			}
		}

		c := 0
		values[14] = values[1]
		for _, val := range values {
			if val != 0 {
				c += 1
			} else {
				c = 0
			}

			if c == 5 {
				return true
			}
		}
	}

	return false
}

func (ch *CheckHand) FourOfAKind(h []deck.Card) bool {
	var values [14]uint16

	for _, c := range h {
		values[c.Value] += 1
	}

	for _, v := range values {
		if v == 4 {
			return true
		}
	}

	return false
}

func (ch *CheckHand) FullHouse(h []deck.Card) bool {
	return ch.ThreeOfAKind(h) && ch.OnePair(h)
}

func (ch *CheckHand) Flush(h []deck.Card) bool {
	var suits [4]uint16

	for _, c := range h {
		suits[c.Suit] += 1
		if suits[c.Suit] == 5 {
			return true
		}
	}

	return false
}

func (ch *CheckHand) Straight(h []deck.Card) bool {
	var values [15]uint16

	for _, c := range h {
		values[c.Value] += 1
	}

	values[14] = values[1]

	c := 0
	for _, val := range values {
		if val != 0 {
			c += 1
		} else {
			c = 0
		}

		if c == 5 {
			return true
		}
	}

	return false
}

func (ch *CheckHand) ThreeOfAKind(h []deck.Card) bool {
	var values [14]uint16

	for _, c := range h {
		values[c.Value] += 1
	}

	for _, v := range values {
		if v == 3 {
			return true
		}
	}

	return false
}

func (ch *CheckHand) TwoPair(h []deck.Card) bool {
	var values [14]uint16

	for _, c := range h {
		values[c.Value] += 1
	}

	pairs := 0
	for _, v := range values {
		if v == 2 {
			pairs += 1
		}
	}

	if pairs >= 2 {
		return true
	}

	return false
}

func (ch *CheckHand) OnePair(h []deck.Card) bool {
	var values [14]uint16

	for _, c := range h {
		values[c.Value] += 1
	}

	pairs := 0
	for _, v := range values {
		if v == 2 {
			pairs += 1
		}
	}

	if pairs == 1 {
		return true
	}

	return false
}
