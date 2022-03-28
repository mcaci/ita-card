package card

import "testing"

func TestCardFromToIDNoErr(t *testing.T) {
	_, err := FromID(1)
	if err != nil {
		t.Fatal("Expecting 1 as result")
	}
}

func TestCardMustIDNoErr(t *testing.T) {
	id := MustID(1)
	if id.Number() != 1 {
		t.Fatal("Expecting 1 as result")
	}
}
