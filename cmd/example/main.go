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

	if input, err := os.Open(*inputFile); err == nil {
		eventLoop := new(engine.EventLoop)
		eventLoop.Start()
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			eventLoop.Post(parse(scanner.Text()))
		}
		eventLoop.AwaitFinish()
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}