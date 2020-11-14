package commands

import (
	"github.com/jn-lp/se-lab4/engine"
)

type reverseCommand struct {
	Arg string
}

func (r *reverseCommand) Execute(loop engine.Handler) {
	var reversed string
	for _, r := range r.Arg {
		reversed = string(r) + reversed
	}
	loop.Post(&printCommand{Arg: reversed})
}
