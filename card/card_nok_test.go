package card

import (
	"testing"
)

var errCheck = func(card *Item, err error) bool { return err == nil }

func Test15OfCupDoesntExist(t *testing.T) {
	applyCheck(t, "15", "Cup", errCheck, "The %s of %s is not a valid card")
}

func Test8OfSpadesDoesntExist(t *testing.T) {
	applyCheck(t, "8", "Spades", errCheck, "The %s of %s is not a valid card")
}

func TestTwoOfCudgelIsIncorrect(t *testing.T) {
	applyCheck(t, "Two", "Cudgel", errCheck, "The %s of %s is not a valid card")
}

func TestEmptyNumberIsIncorrect(t *testing.T) {
	applyCheck(t, "", "Cudgel", errCheck, "The %s of %s is not a valid card")
}

func TestEmptySeedIsIncorrect(t *testing.T) {
	applyCheck(t, "6", "", errCheck, "The %s of %s is not a valid card")
}
