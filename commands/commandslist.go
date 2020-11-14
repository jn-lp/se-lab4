package commands

import (
	"reflect"
	"strconv"

	"github.com/jn-lp/se-lab4/engine"
)

var commandsList = map[string]engine.Command{
	"print": &printCommand{},
	"mul": &mulCommand{},
	"printc": &printcCommand{},
	"add": &addCommand{},
	"reverse": &reverseCommand{},
	"palindrome": &palindromeCommand{},
	"sha1": &sha1Command{},
	"split": &splitCommand{},
	"delete": &deleteCommand{},
	"cat": &catCommand{},
}

func Construct(commandName string, args []string) (command engine.Command) {
	command = commandsList[commandName]
	commandReflection := reflect.ValueOf(command)
	if len(args) != commandReflection.Elem().NumField() {
		command = &printCommand{Arg: "The number of args is out of index"}
		return
	}

	for k, arg := range args {
		switch field := commandReflection.Elem().Field(k); field.Type().Kind() {
		case reflect.Int:
			intArg, err := strconv.Atoi(arg)
			if err != nil {
				command = &printCommand{Arg: "SYNTAX ERROR: " + err.Error()}
			}
			field.SetInt(int64(intArg))
		default:
			field.SetString(arg)
		}
	}

	return
}
