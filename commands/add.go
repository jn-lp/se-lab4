package commands

import (
	"strconv"

	"github.com/jn-lp/se-lab4/engine"
)

type addCommand struct {
	Arg1, Arg2 int
}

func (a *addCommand) Execute(loop engine.Handler) {
	sum := a.Arg1 + a.Arg2
	loop.Post(&printCommand{Arg: strconv.Itoa(sum)})
}
