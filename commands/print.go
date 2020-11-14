package commands

import (
	"fmt"

	"github.com/jn-lp/se-lab4/engine"
)

type printCommand struct {
	Arg string
}

func (p *printCommand) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}
