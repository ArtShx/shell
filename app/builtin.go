package main

import (
	"fmt"
	"os"
	"strings"
)

type BuiltinCommands interface {
	Run()
}

type ExitCommand struct{}
type EchoCommand struct {
	Args []string
}
type TypeCommand struct {
	Args []string
}

func (c ExitCommand) Run() {
	os.Exit(0)
}

func (c EchoCommand) Run() {
	fmt.Println(strings.Join(c.Args[1:], " "))
}

func (c TypeCommand) Run() {
	if len(c.Args) <= 1 {
		return
	}
	_, ok := ParseBuiltinCommands(c.Args[1:])
	if ok == nil {
		fmt.Printf("%s is a shell builtin\n", c.Args[1])
	} else {
		fmt.Printf("%s: not found\n", c.Args[1])
	}
}

const (
	exit  string = "exit"
	echo  string = "echo"
	type_ string = "type"
)

func ParseBuiltinCommands(args []string) (bc BuiltinCommands, err error) {
	switch args[0] {
	case "exit":
		return ExitCommand{}, nil
	case "echo":
		return EchoCommand{args}, nil
	case "type":
		return TypeCommand{args}, nil
	default:
		return nil, fmt.Errorf("unknown command [%s]", args)
	}
}
