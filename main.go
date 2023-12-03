package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/HarshitCd/GoPoker/deck"
)

func main() {
	rand.Seed(time.Now().Unix())

	myDeck := deck.New()
	fmt.Println(myDeck)
}
