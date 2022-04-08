package card_test

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func TestCardFromToIDNoErr(t *testing.T) {
	testcases := []struct {
		name         string
		id           int
		expectErr    bool
		expectedCard *card.Item
	}{
		{"Card created correctly", 1, false, card.MustID(1)},
		{"Id out of bounds for card id", 100, true, nil},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			item, err := card.FromID(tc.id)
			switch tc.expectErr {
			case true:
				if err != nil {
					break
				}
				t.Fatal("Expecting an error but got nil")
			case false:
				if err != nil {
					t.Fatalf("Not expecting any errors but got %v", err)
				}
				if item.Number() != tc.expectedCard.Number() {
					t.Fatalf("Expecting number %v but got %v", tc.expectedCard.Number(), item.Number())
				}
				if item.Seed() != tc.expectedCard.Seed() {
					t.Fatalf("Expecting seed %v but got %v", tc.expectedCard.Seed(), item.Seed())
				}
			}
		})
	}
}
