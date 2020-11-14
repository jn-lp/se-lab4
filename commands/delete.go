package commands

import (
	"strings"

	"github.com/jn-lp/se-lab4/engine"
)

type deleteCommand struct {
	Str, Symbol string
}

func (d *deleteCommand) Execute(loop engine.Handler) {
	deleted := strings.ReplaceAll(d.Str, d.Symbol, "")
	loop.Post(&printCommand{Arg: deleted})
}
