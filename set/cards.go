package set

import "github.com/mcaci/ita-cards/card"

// Cards type
type Cards []card.Item

// Add func
func (cards *Cards) Add(ids ...card.Item) {
	*cards = append(*cards, ids...)
}

// Clear func
func (cards *Cards) Clear() {
	*cards = Cards{}
}

// Sum func
func (cards *Cards) Sum(point func(card.Item) uint8) (sum uint8) {
	for _, c := range *cards {
		sum += point(c)
	}
	return
}

// Find func
func (cards *Cards) Find(id card.Item) int {
	for index, c := range *cards {
		if c == id {
			return index
		}
	}
	return -1
}

// Supply func
func (cards *Cards) Supply() card.Item {
	card := (*cards)[0]
	*cards = (*cards)[1:]
	return card
}
