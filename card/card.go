// Package card contains all the utilities to manipulate italian-style cards (set of 40 cards)
package card

import (
	"fmt"
)

// Item represents a card with number and seed
type Item struct {
	number uint8
	seed   Seed
}

// Number returns the number value of the card (1 to 10)
func (c Item) Number() uint8 {
	return c.number
}

// Seed returns the seed value of the card as Seed type (Coin, Cup, Sword, Cudgel)
func (c Item) Seed() Seed {
	return c.seed
}

func (c Item) String() string {
	if c.Number() == 0 && c.Seed() == 0 {
		return "(Undefined card)"
	}
	return fmt.Sprintf("%d%s", c.Number(), c.Seed())
}
