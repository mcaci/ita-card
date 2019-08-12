package bench

import (
	"fmt"
)

func origCreate(number, seed string) (id, error) {
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

func oldCreate(number, seed string) (id, error) {
	s, errS := toS(seed)
	n, errN := toN(number)
	var err error
	if errN != nil || errS != nil {
		err = fmt.Errorf("%v; %v", errN, errS)
	}
	return id(n + s), err
}

func seqCreate(number, seed string) (id, error) {
	var c cardWithErr
	c.eval(toS, seed)
	c.eval(toN, number)
	return id(c.id), c.err
}

func concurCreate(number, seed string) (id, error) {
	var c cardWithErr
	done := make(chan bool, 2)
	defer close(done)
	go func() { c.eval(toN, number); done <- true }()
	go func() { c.eval(toS, seed); done <- true }()
	<-done
	<-done
	return id(c.id), c.err
}
