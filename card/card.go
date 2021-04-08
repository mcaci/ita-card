package card

import "fmt"

// Item represents a card with number and seed
type Item struct {
	number uint8
	seed   Seed
}

// New creates a new card Item (a card with number and seed)
func New(number, seed string) (*Item, error) {
	c, err := create(number, seed)
	item := Item{number: c.number(), seed: c.seed()}
	return &item, err
}

// FromID creates a new card Item (a card with number and seed)
// from an id ranging from 1 to 40
func FromID(n uint8) (*Item, error) {
	if n < 1 || n > 40 {
		return nil, fmt.Errorf("%d is not between the inclusive range [1,40]", n)
	}
	return &Item{number: id(n).number(), seed: id(n).seed()}, nil
}

// MustID creates a new card Item (a card with number and seed)
// from an id ranging from 1 to 40, panics if outside
func MustID(n uint8) *Item {
	if n < 1 || n > 40 {
		panic(fmt.Errorf("%d is not between the inclusive range [1,40]", n))
	}
	return &Item{number: id(n).number(), seed: id(n).seed()}
}

// Number returns the number value of the card (1 to 10)
func (c Item) Number() uint8 {
	return c.number
}

// Seed returns the seed value of the card as Seed type (Coin, Cup, Sword, Cudgel)
func (c Item) Seed() Seed {
	return c.seed
}

// ToID returns the id value of the card (1 to 40)
// 1-10 are Coin cards
// 11-20 are Cup cards
// 21-30 are Sword cards
// 31-40 are Cudgel cards
func (c Item) ToID() uint8 {
	return c.number + uint8(c.seed)*10
}

func (c Item) String() string {
	if c.Number() == 0 && c.Seed() == 0 {
		return "(Undefined card)"
	}
	return fmt.Sprintf("%d%s", c.Number(), c.Seed())
}
