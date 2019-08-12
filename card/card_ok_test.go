package card

import (
	"strconv"
	"testing"
)

var noErrCheck = func(card *Item, err error) bool { return err != nil }

func applyCheck(t *testing.T, number, seed string, f func(card *Item, err error) bool, message string) {
	if f(New(number, seed)) {
		t.Fatalf(message, number, seed)
	}
}

func Test1OfCoinIsCreatedCorrectly_NoError(t *testing.T) {
	applyCheck(t, "1", "Coin", noErrCheck, "An unexpected error was raised")
}

func Test1OfCoinIsCreatedCorrectly_NumberIs1(t *testing.T) {
	number := "1"
	check := func(card *Item, err error) bool { return strconv.Itoa(int(card.Number())) != number }
	applyCheck(t, number, "Coin", check, "Card's number is not created well from %s and %s")
}

func Test1OfCoinIsCreatedCorrectly_SeedIsCoin(t *testing.T) {
	seed := Coin
	check := func(card *Item, err error) bool { return card.Seed() != seed }
	applyCheck(t, "1", "Coin", check, "Card's seed is not created well from %s and %s")
}

func Test2OfSwordIsCreatedCorrectly_NoError(t *testing.T) {
	applyCheck(t, "2", "Sword", noErrCheck, "An unexpected error was raised")
}

func Test2OfSwordIsCreatedCorrectly_NumberIs2(t *testing.T) {
	number := "2"
	check := func(card *Item, err error) bool { return strconv.Itoa(int(card.Number())) != number }
	applyCheck(t, number, "Sword", check, "Card's number is not created well from %s and %s")
}

func Test2OfSwordIsCreatedCorrectly_SeedIsSword(t *testing.T) {
	seed := Sword
	check := func(card *Item, err error) bool { return card.Seed() != seed }
	applyCheck(t, "2", "Sword", check, "Card's seed is not created well from %s and %s")
}

func Test8OfCupIsCreatedCorrectly_NoError(t *testing.T) {
	applyCheck(t, "8", "Cup", noErrCheck, "An unexpected error was raised")
}

func Test8OfCupIsCreatedCorrectly_NumberIs8(t *testing.T) {
	number := "8"
	check := func(card *Item, err error) bool { return strconv.Itoa(int(card.Number())) != number }
	applyCheck(t, number, "Cup", check, "Card's number is not created well from %s and %s")
}

func Test8OfCupIsCreatedCorrectly_SeedIsCup(t *testing.T) {
	seed := Cup
	check := func(card *Item, err error) bool { return card.Seed() != seed }
	applyCheck(t, "8", "Cup", check, "Card's seed is not created well from %s and %s")
}

func Test10OfCudgelIsCreatedCorrectly_NoError(t *testing.T) {
	applyCheck(t, "10", "Cudgel", noErrCheck, "An unexpected error was raised")
}

func Test10OfCudgelIsCreatedCorrectly_NumberIs10(t *testing.T) {
	number := "10"
	check := func(card *Item, err error) bool { return strconv.Itoa(int(card.Number())) != number }
	applyCheck(t, number, "Cudgel", check, "Card's number is not created well from %s and %s")
}

func Test10OfCudgelIsCreatedCorrectly_SeedIsCudgel(t *testing.T) {
	seed := Cudgel
	check := func(card *Item, err error) bool { return card.Seed() != seed }
	applyCheck(t, "10", "Cudgel", check, "Card's seed is not created well from %s and %s")
}
