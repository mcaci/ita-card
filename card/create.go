package card

import (
	"fmt"
	"strconv"
)

func create(number, seed string) (id, error) {
	n, err := toN(number)
	if err != nil {
		return 0, err
	}
	s, err := toS(seed)
	if err != nil {
		return 0, err
	}
	return id(n + s), err
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
		return uint8(i) * 10, nil
	}
	return 0, fmt.Errorf("Seed %q doesn't exist", seed)
}
