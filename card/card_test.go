package card

import (
	"strconv"
	"testing"
)

func TestCardMustToID(t *testing.T) {
	if MustID(1).ToID() != 1 {
		t.Fatal("Expecting 1 as result")
	}
}

func TestCardFromToIDNoErr(t *testing.T) {
	_, err := FromID(1)
	if err != nil {
		t.Fatal("Expecting 1 as result")
	}
}

func TestCardCreation(t *testing.T) {
	tcs := []struct {
		name   string
		number string
		seed   string
		msg    string
		errF   func(err error) bool
		nCheck func(card *Item, number string) bool
		sCheck func(card *Item, seed string) bool
	}{
		{"1OfCoinIsCreatedCorrectly_NoError", "1", "Coin", "An unexpected error was raised", noErrCheck, nil, nil},
		{"1OfCoinIsCreatedCorrectly_NumberIs1", "1", "Coin", "Card's number is not created well from %q and %q", nil, nCheck, nil},
		{"Test1OfCoinIsCreatedCorrectly_SeedIsCoin", "1", "Coin", "Card's seed is not created well from %q and %q", nil, nil, sCheck},
		{"Test2OfSwordIsCreatedCorrectly_NoError", "2", "Sword", "An unexpected error was raised", noErrCheck, nil, nil},
		{"Test2OfSwordIsCreatedCorrectly_NumberIs2", "2", "Sword", "Card's number is not created well from %q and %q", nil, nCheck, nil},
		{"Test2OfSwordIsCreatedCorrectly_SeedIsSword", "2", "Sword", "Card's seed is not created well from %q and %q", nil, nil, sCheck},
		{"Test8OfCupIsCreatedCorrectly_NoError", "8", "Cup", "An unexpected error was raised", noErrCheck, nil, nil},
		{"Test8OfCupIsCreatedCorrectly_NumberIs8", "8", "Cup", "Card's number is not created well from %q and %q", nil, nCheck, nil},
		{"Test8OfCupIsCreatedCorrectly_SeedIsCup", "8", "Cup", "Card's seed is not created well from %q and %q", nil, nil, sCheck},
		{"Test10OfCudgelIsCreatedCorrectly_NoError", "10", "Cudgel", "An unexpected error was raised", noErrCheck, nil, nil},
		{"Test10OfCudgelIsCreatedCorrectly_NumberIs10", "10", "Cudgel", "Card's number is not created well from %q and %q", nil, nCheck, nil},
		{"Test10OfCudgelIsCreatedCorrectly_SeedIsCudgel", "10", "Cudgel", "Card's seed is not created well from %q and %q", nil, nil, sCheck},
		{"Test15OfCupDoesntExist", "15", "Cup", "The %q of %q is not a valid card", errCheck, nil, nil},
		{"Test8OfSpadesDoesntExist", "8", "Spades", "The %q of %q is not a valid card", errCheck, nil, nil},
		{"TestTwoOfCudgelIsIncorrect", "Two", "Cudgel", "The %q of %q is not a valid card", errCheck, nil, nil},
		{"TestEmptyNumberIsIncorrect", "", "Cudgel", "The %q of %q is not a valid card", errCheck, nil, nil},
		{"TestEmptySeedIsIncorrect", "6", "", "The %q of %q is not a valid card", errCheck, nil, nil},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			item, err := New(tc.number, tc.seed)
			switch {
			case tc.errF != nil:
				if tc.errF(err) {
					t.Fatalf(tc.msg, tc.number, tc.seed)
				}
			case tc.nCheck != nil:
				if tc.nCheck(item, tc.number) {
					t.Fatalf(tc.msg, tc.number, tc.seed)
				}
			case tc.sCheck != nil:
				if tc.sCheck(item, tc.seed) {
					t.Fatalf(tc.msg, tc.number, tc.seed)
				}
			}
		})
	}
}

var errCheck = func(err error) bool { return err == nil }
var noErrCheck = func(err error) bool { return err != nil }
var nCheck = func(card *Item, number string) bool { return strconv.Itoa(int(card.Number())) != number }
var sCheck = func(card *Item, seed string) bool { return card.Seed().String() != string(seed) }
