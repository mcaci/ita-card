package card

//go:generate stringer -type Seed seed.go
// generated using http://godoc.org/golang.org/x/tools/cmd/stringer
// Seed type is an id for the seed ranging from 0 to 3
type Seed uint8

const (
	// Coin is the english equivalent of ORO
	Coin Seed = iota
	// Cup is the english equivalent of COPPE
	Cup
	// Sword is the english equivalent of SPADE
	Sword
	// Cudgel is the english equivalent of BASTONI/MAZZE
	Cudgel
)

func (s Seed) Seed() Seed { return s }
