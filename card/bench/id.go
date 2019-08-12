// Package bench only used for benchmarking card creation
package bench

import "github.com/mcaci/ita-cards/card"

// id is the id of a card from 1 to 40
type id uint8

func (c id) number() uint8 {
	return (uint8(c)-1)%10 + 1
}

func (c id) seed() card.Seed {
	return card.Seed((uint8(c) - 1) / 10)
}
