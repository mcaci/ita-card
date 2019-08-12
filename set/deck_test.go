package set

import "testing"

func TestCreateDeck(t *testing.T) {
	d := Deck()
	if len(d) != DeckSize {
		t.Fatalf("There should be 40 card in the deck")
	}
}

func TestRemovingTwoCardsShouldGiveDifferentCards(t *testing.T) {
	d := Deck()
	a := d.Top()
	b := d.Top()

	if a == b {
		t.Fatalf("Drawn cards should be different but they are %v and %v", a, b)
	}
}
