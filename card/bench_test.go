package card_test

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func BenchmarkCreationViaNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		card.New("1", "Coin")
	}
}

func BenchmarkCreationViaID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		card.FromID(1)
	}
}

func BenchmarkCreationViaMustID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		card.MustID(1)
	}
}
