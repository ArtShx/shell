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

func (c ExitCommand) Run() {
	os.Exit(0)
}

func (c EchoCommand) Run() {
	fmt.Println(strings.Join(c.Args[1:], " "))
}

const (
	exit string = "exit"
	echo string = "echo"
)

func ParseBuiltinCommands(args []string) (bc BuiltinCommands, err error) {
	switch args[0] {
	case "exit":
		return ExitCommand{}, nil
	case "echo":
		return EchoCommand{args}, nil
	default:
		return nil, fmt.Errorf("unknown command [%s]", args)
	}
}
