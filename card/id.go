// Package card contains all the utilities to manipulate italian-style cards (set of 40 cards)
package card

// id is the id of a card from 1 to 40
type id uint8

func (c id) number() uint8 {
	return (uint8(c)-1)%10 + 1
}

func (c id) seed() Seed {
	return Seed((uint8(c) - 1) / 10)
}
