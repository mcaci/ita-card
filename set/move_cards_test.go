package set

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"golang.org/x/exp/constraints"
)

// toID returns the id value of the card (1 to 40)
// 1-10 are Coin cards
// 11-20 are Cup cards
// 21-30 are Sword cards
// 31-40 are Cudgel cards
func toID[T constraints.Integer](c card.Item) T {
	return T(c.Number() + uint8(c.Seed())*10)
}

func TestCardMustToID(t *testing.T) {
	if toID[int](*card.MustID(1)) != 1 {
		t.Fatal("Expecting 1 as result")
	}
}

func TestCardsAreMoved(t *testing.T) {
	from := NewMust(1, 2, 3, 4, 5)
	to := NewMust[uint8]()
	Move(from, to)
	if len(*from) > 0 {
		t.Fatalf("from should be empty but has %v", from)
	}
}

func TestPlayerCountPoints(t *testing.T) {
	from := NewMust(1, 2, 3, 4, 5)
	to := NewMust[uint8]()
	Move(from, to)
	if len(*to) == 0 {
		t.Fatal("to should not be empty but is")
	}
}

func TestMoveOneCard(t *testing.T) {
	from := NewMust(1, 2, 3, 4, 5)
	to := NewMust[int]()
	MoveOne(card.MustID(2), from, to)
	if len(*to) == 0 {
		t.Fatal("to should not be empty but is")
	}
	if id := toID[int](*card.MustID(1)); id != 1 {
		t.Fatalf("moved card should be the one with ID == 2 but was %d", id)
	}
}

func TestCardMovedCorrectlyFromSource(t *testing.T) {
	from := NewMust(1, 2, 3, 4, 5)
	lenFrom := len(*from)
	to := NewMust[uint8]()
	MoveOne(card.MustID(2), from, to)
	if lenFrom == len(*from) {
		t.Fatalf("From should be 4 but was %d: from %v, to %v", lenFrom, from, to)
	}
	if id := toID[int](*card.MustID(1)); id != 1 {
		t.Fatalf("moved card should be the one with ID == 2 but was %d", id)
	}
}

func TestMoveOneCardError(t *testing.T) {
	from := NewMust(1, 2, 3, 4, 5)
	to := NewMust[uint8]()
	err := MoveOne(card.MustID(21), from, to)
	if err == nil {
		t.Fatal("expecting a not found error here but no errors were raised")
	}
}
