package set

import "github.com/mcaci/ita-cards/card"

// Cards type
type Cards []card.Item

// NewMust creates a set of cards with ids that must be valid
// (from 1 to 40) or else it will panic
func NewMust(ids ...uint8) *Cards {
	cards := make(Cards, len(ids))
	for i, id := range ids {
		cards[i] = *card.MustID(id)
	}
	return &cards
}

// Add appends cards to the set
func (cards *Cards) Add(ids ...card.Item) {
	*cards = append(*cards, ids...)
}

// Clear removes all cards from the set
func (cards *Cards) Clear() {
	*cards = Cards{}
}

// Find returns the index of the card given in input or
// -1 if not found (to change with error)
func (cards *Cards) Find(id card.Item) int {
	for index, c := range *cards {
		if c == id {
			return index
		}
	}
	return -1
}

// Top func return the top card of the set
// and deletes from it
func (cards *Cards) Top() card.Item {
	card := (*cards)[0]
	*cards = (*cards)[1:]
	return card
}
