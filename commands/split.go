package commands

import (
	"strings"

	"github.com/jn-lp/se-lab4/engine"
)

type splitCommand struct {
	Str, Sep string
}

func (s *splitCommand) Execute(loop engine.Handler) {
	splitted := strings.Split(s.Str, s.Sep)

	for _, s := range splitted {
		loop.Post(&printCommand{Arg: s})
	}
}
