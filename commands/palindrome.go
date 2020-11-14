package commands

import (
	"github.com/jn-lp/se-lab4/engine"
)

type palindromeCommand struct {
	Arg string
}

func (p *palindromeCommand) Execute(loop engine.Handler) {
	var reversed string
	for _, r := range p.Arg {
		reversed = string(r) + reversed
	}
	palindrom := p.Arg + reversed
	loop.Post(&printCommand{Arg: palindrom})
}
