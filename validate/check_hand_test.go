package validate

import (
	"testing"

	"github.com/HarshitCd/GoPoker/deck"
)

func TestStraightFlush(t *testing.T) {

	h := []deck.Card{
		deck.NewCard(0, 1),
		deck.NewCard(1, 2),
		deck.NewCard(0, 3),
		deck.NewCard(0, 4),
		deck.NewCard(0, 5),
		deck.NewCard(3, 8),
		deck.NewCard(2, 10),
	}

	ch := New()
	actual := ch.StraightFlush(h)
	expected := true

	if expected != actual {
		t.Errorf("Expected a %v but got %v", expected, actual)
	}
}

func TestStraight(t *testing.T) {

	h := []deck.Card{
		deck.NewCard(0, 1),
		deck.NewCard(1, 2),
		deck.NewCard(0, 3),
		deck.NewCard(0, 4),
		deck.NewCard(0, 5),
		deck.NewCard(3, 8),
		deck.NewCard(2, 10),
	}

	ch := New()
	actual := ch.Straight(h)
	expected := true

	if expected != actual {
		t.Errorf("Expected a %v but got %v", expected, actual)
	}
}
