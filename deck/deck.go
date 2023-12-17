package deck

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

func (s Suit) String() string {
	switch s {
	case 0:
		return "SPADES"
	case 1:
		return "HEARTS"
	case 2:
		return "DIAMONDS"
	case 3:
		return "CLUBS"
	default:
		panic("Invalid suit...")
	}
}

func suitToUniCode(s Suit) string {
	switch s {
	case 0:
		return "♠"
	case 1:
		return "♥"
	case 2:
		return "♦"
	case 3:
		return "♣"
	default:
		panic("Invalid suit...")
	}
}

type Card struct {
	Suit  Suit
	Value int
}

func (c Card) String() string {

	var value string

	switch c.Value {
	case 1:
		value = "ACE"
	case 11:
		value = "JACK"
	case 12:
		value = "QUEEN"
	case 13:
		value = "KING"
	default:
		value = strconv.Itoa(c.Value)
	}

	return fmt.Sprintf("%s %s", value, suitToUniCode(c.Suit))
}

func NewCard(s Suit, v int) Card {

	if v > 13 {
		panic("Value of the cards cannot be higher than 13.")
	}

	return Card{
		Suit:  s,
		Value: v,
	}
}

type Deck [52]Card

func New() Deck {
	var (
		newDeck = Deck{}
		nCards  = 13
		nSuits  = 4
	)

	for i := 1; i <= nCards; i++ {
		for j := 0; j < nSuits; j++ {
			newDeck[j*nCards+i-1] = NewCard(Suit(j), i)
		}
	}

	return Shuffle(newDeck)
}

func Shuffle(d Deck) Deck {

	for i := 0; i < 52; i++ {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}

	return d
}
