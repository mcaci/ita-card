package set

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func TestCardsAreMoved(t *testing.T) {
	from := NewMust(1, 2, 3, 4, 5)
	to := NewMust()
	Move(from, to)
	if len(*from) > 0 {
		t.Fatalf("from should be empty but has %v", from)
	}
}

func TestPlayerCountPoints(t *testing.T) {
	from := NewMust(1, 2, 3, 4, 5)
	to := NewMust()
	Move(from, to)
	if len(*to) == 0 {
		t.Fatal("to should not be empty but is")
	}
}

func TestMoveOneCard(t *testing.T) {
	from := NewMust(1, 2, 3, 4, 5)
	to := NewMust()
	MoveOne(card.MustID(2), from, to)
	if len(*to) == 0 {
		t.Fatal("to should not be empty but is")
	}
	if id := (*to)[0].ToID(); id != 2 {
		t.Fatalf("moved card should be the one with ID == 2 but was %d", id)
	}
}

func TestCardMovedCorrectlyFromSource(t *testing.T) {
	from := NewMust(1, 2, 3, 4, 5)
	lenFrom := len(*from)
	to := NewMust()
	MoveOne(card.MustID(2), from, to)
	if lenFrom != len(*from) {
		t.Fatalf("From should be 4 but was %d: from %v, to %v", lenFrom, from, to)
	}
	if id := (*to)[0].ToID(); id != 2 {
		t.Fatalf("moved card should be the one with ID == 2 but was %d", id)
	}
}

func TestMoveOneCardError(t *testing.T) {
	from := NewMust(1, 2, 3, 4, 5)
	to := NewMust()
	err := MoveOne(card.MustID(21), from, to)
	if err == nil {
		t.Fatal("expecting a not found error here but no errors were raised")
	}
}
