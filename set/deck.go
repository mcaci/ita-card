package set

import (
	"math/rand"
	"time"

	"github.com/mcaci/ita-cards/card"
)

// DeckSize of a cards of cards
const DeckSize = 40

// Deck func
func Deck() (cards Cards) {
	rand.Seed(time.Now().UnixNano())
	ints := rand.Perm(DeckSize)
	for _, cardID := range ints {
		c := card.MustID(fromZeroBased(cardID))
		cards.Add(*c)
	}
	return
}

func fromZeroBased(index int) uint8 {
	return uint8(index) + 1
}
