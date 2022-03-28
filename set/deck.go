package set

import (
	"math/rand"
	"time"

	"github.com/mcaci/ita-cards/card"
)

// DeckSize of a cards of cards
const DeckSize = 40

// Deck returns a full shuffled deck with all the 40 cards
func Deck() (cards Cards) {
	rand.Seed(time.Now().UnixNano())
	ints := rand.Perm(DeckSize)
	for _, cardID := range ints {
		c := card.MustID(cardID + 1)
		cards.Add(*c)
	}
	return
}
