package main

import (
	"fmt"

	"github.com/HarshitCd/GoPoker/deck"
	"github.com/HarshitCd/GoPoker/validate"
)

func main() {
	myDeck := deck.New()

	noOfPlayers := 10
	var hands [][]deck.Card
	for i := 0; i < noOfPlayers; i++ {
		h := make([]deck.Card, 2)
		copy(h, myDeck[2*i:2*i+2])

		hands = append(hands, h)
	}
	flop := make([]deck.Card, 5)
	copy(flop, myDeck[2*noOfPlayers+2:2*noOfPlayers+7])

	chechHand := validate.New()
	for i, h := range hands {
		fmt.Println("Flop:", flop)
		fmt.Printf("Hand %d: %v\n", i+1, h)
		chechHand.Validate(h, flop)
		fmt.Println("----------------------------------")
	}

}
