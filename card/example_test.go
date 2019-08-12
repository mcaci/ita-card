package card_test

import (
	"fmt"

	"github.com/mcaci/ita-cards/card"
)

func ExampleNew() {
	c, err := card.New("1", "Coin")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}
