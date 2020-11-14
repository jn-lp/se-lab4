package commands

import (
	"github.com/jn-lp/se-lab4/engine"
)

type catCommand struct {
	Arg1, Arg2 string
}

func (c *catCommand) Execute(loop engine.Handler) {
	concatenated := c.Arg1 + c.Arg2
	loop.Post(&printCommand{Arg: concatenated})
}
