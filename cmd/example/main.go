package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/jn-lp/se-lab4/commands"
	"github.com/jn-lp/se-lab4/engine"
)

func parse(commandline string) engine.Command {
	commandFields := strings.Fields(commandline)
	commandName, args := commandFields[0], commandFields[1:]
	command := commands.Construct(commandName, args)
	return command
}

func main() {
	inputFile := flag.String("f", "", "File to run with EventLoop")
	flag.Parse()

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(*inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			eventLoop.Post(parse(scanner.Text()))
		}
	}
	eventLoop.AwaitFinish()
}
