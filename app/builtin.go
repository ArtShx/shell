package main

import (
	"fmt"
	"os"
	"strings"
)

type BaseCommand struct {
	Args []string
}

func (b BaseCommand) GetName() string {
	fmt.Println("GetName", b.Args)
	if len(b.Args) == 0 {
		return ""
	}
	return b.Args[0]
}

type BuiltinCommands interface {
	Run()
	GetName() string
}

type ExitCommand struct {
	BaseCommand
}
type EchoCommand struct {
	BaseCommand
}
type TypeCommand struct {
	BaseCommand
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
	base := BaseCommand{Args: args}
	switch args[0] {
	case "exit":
		return ExitCommand{base}, nil
	case "echo":
		return EchoCommand{base}, nil
	case "type":
		return TypeCommand{base}, nil
	default:
		return nil, fmt.Errorf("unknown command [%s]", args)
	}
}
