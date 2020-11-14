package commands

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/jn-lp/se-lab4/engine"
)

type sha1Command struct {
	Arg string
}

func (s *sha1Command) Execute(loop engine.Handler) {
	h := sha1.New()
	h.Write([]byte(s.Arg))

	hex := hex.EncodeToString(h.Sum(nil))
	loop.Post(&printCommand{Arg: hex})
}
