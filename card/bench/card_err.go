package bench

import (
	"fmt"
)

type cardWithErr struct {
	id  uint8
	err error
}

func (c *cardWithErr) eval(f accumulator, val string) {
	id, err := f(val)
	if c.err == nil && err == nil {
		c.id += id
		return
	}
	if c.err == nil {
		c.err = err
		return
	}
	c.err = fmt.Errorf("%v; %v", c.err, err)
}
