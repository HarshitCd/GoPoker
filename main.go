package main

import (
	"fmt"

	"github.com/HarshitCd/GoPoker/deck"
	"github.com/HarshitCd/GoPoker/validate"
)

func main() {
	myDeck := deck.New()

	h := myDeck[:2]
	flop := myDeck[2:7]
	k := make([]deck.Card, len(h))
	copy(k, h)

	fmt.Println(h, flop)
	checkHand := validate.New()
	checkHand.Validate(h, flop)
}
