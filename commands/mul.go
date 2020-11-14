package commands

import (
	"strconv"

	"github.com/jn-lp/se-lab4/engine"
)

type mulCommand struct {
	Arg1, Arg2 int
}

func (m *mulCommand) Execute(loop engine.Handler) {
	product := m.Arg1 * m.Arg2
	loop.Post(&printCommand{Arg: strconv.Itoa(product)})
}
