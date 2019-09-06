package card

import "testing"

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
