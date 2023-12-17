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

	fmt.Println("Straight Flush:", ch.straightFlush(f))
	fmt.Println("Four Of A Kind:", ch.fourOfAKind(f))
	fmt.Println("Full House:", ch.fullHouse(f))
	fmt.Println("Flush:", ch.flush(f))
	fmt.Println("Straight:", ch.straight(f))
	fmt.Println("Three Of A Kind:", ch.threeOfAKind(f))
	fmt.Println("Two Pair:", ch.twoPair(f))
	fmt.Println("One Pair:", ch.onePair(f))
}

func (ch *CheckHand) straightFlush(h []deck.Card) bool {
	var values [15]uint16

	if h[0].Suit != h[1].Suit {
		return false
	}

	suit := h[0].Suit

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
			return h[0].Value != h[1].Value
		}
	}

	return false
}

func (ch *CheckHand) fourOfAKind(h []deck.Card) bool {
	var values [14]uint16

	for _, c := range h {
		values[c.Value] += 1
	}

	for _, val := range values {
		if val == 4 {
			return true
		}
	}

	return false
}

func (ch *CheckHand) fullHouse(h []deck.Card) bool {
	var suits [14]uint16

	for _, c := range h {
		suits[c.Value] += 1
	}

	return (suits[h[0].Value] == 3 && suits[h[1].Value] == 2) || (suits[h[0].Value] == 2 && suits[h[1].Value] == 3)
}

func (ch *CheckHand) flush(h []deck.Card) bool {
	var suits [4]uint16

	for _, c := range h {
		suits[c.Suit] += 1
	}

	return h[0].Suit == h[1].Suit && suits[h[0].Suit] == 5
}

func (ch *CheckHand) straight(h []deck.Card) bool {
	var values [15]uint16

	if h[0].Value == h[1].Value {
		return false
	}

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
			return values[h[0].Value] != 0 && values[h[1].Value] != 0
		}
	}

	return false
}

func (ch *CheckHand) threeOfAKind(h []deck.Card) bool {
	var values [14]uint16

	for _, c := range h {
		values[c.Value] += 1
	}

	return values[h[0].Value] == 3 || values[h[1].Value] == 3
}

func (ch *CheckHand) twoPair(h []deck.Card) bool {
	var values [14]uint16

	for _, c := range h {
		values[c.Value] += 1
	}

	return h[0].Value != h[1].Value && values[h[0].Value] == 2 && values[h[1].Value] == 2
}

func (ch *CheckHand) onePair(h []deck.Card) bool {
	var values [14]uint16

	for _, c := range h {
		values[c.Value] += 1
	}

	return values[h[0].Value] == 2 || values[h[1].Value] == 2
}
