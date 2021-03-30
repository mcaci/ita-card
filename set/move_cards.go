package set

import (
	"fmt"

	"github.com/mcaci/ita-cards/card"
)

// Move moves all cards from first to second set in parameters
func Move(from, to *Cards) {
	to.Add(*from...)
	from.Clear()
}

// MoveOne moves one card from first to second set in parameters
func MoveOne(c *card.Item, from, to *Cards) error {
	i := from.Find(*c)
	if i < 0 {
		return fmt.Errorf("card (%v) not found in the source set (%v)", *c, *from)
	}
	to.Add((*from)[i])
	*from = append((*from)[:i], (*from)[i+1:]...)
	return nil
}
