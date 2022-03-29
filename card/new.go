package card

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

// FromID creates a new card Item (a card with number and seed)
// from an id ranging from 1 to 40
// 1-10 are Coin cards
// 11-20 are Cup cards
// 21-30 are Sword cards
// 31-40 are Cudgel cards
func FromID[T constraints.Integer](id T) (*Item, error) {
	if id < 1 || id > 40 {
		return nil, fmt.Errorf("%d is not between the inclusive range [1,40]", id)
	}
	number := func(n T) uint8 { return uint8((n-1)%10 + 1) }
	seed := func(n T) Seed { return Seed((n - 1) / 10) }
	return &Item{number: number(id), seed: seed(id)}, nil
}

// MustID creates a new card Item (a card with number and seed)
// from an id ranging from 1 to 40, panics if outside
func MustID[T constraints.Integer](n T) *Item {
	id, err := FromID(n)
	if err != nil {
		panic(err)
	}
	return id
}

// New creates a new card Item (a card with number and seed)
func New(number, seed string) (*Item, error) {
	n, err := toN(number)
	if err != nil {
		return nil, err
	}
	s, err := toS(seed)
	if err != nil {
		return nil, err
	}
	item := Item{number: n, seed: Seed(s)}
	return &item, nil
}

func toN(number string) (uint8, error) {
	n, err := strconv.Atoi(number)
	if n > 10 || n < 1 {
		err = fmt.Errorf("Number %q is not valid for card", number)
	}
	return uint8(n), err
}

func toS(seed string) (uint8, error) {
	seeds := []string{"Coin", "Cup", "Sword", "Cudgel"}
	for i, s := range seeds {
		if s != seed {
			continue
		}
		return uint8(i), nil
	}
	return 0, fmt.Errorf("Seed %q doesn't exist", seed)
}
