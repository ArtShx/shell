package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type IBaseCommand interface {
	Run()
	GetName() string
	GetGroup() CommandGroup
	GetPath() string
}

type BaseCommand struct {
	IBaseCommand
	Args []string
}

type CommandGroup string

const (
	GroupBuiltin  = "builtin"
	GroupExternal = "external"
)

type BuiltinCommand struct {
	BaseCommand
}

type ExternalCommand struct {
	BaseCommand
	fullPath string
}

func (b BaseCommand) GetName() string {
	fmt.Println("GetName", b.Args)
	if len(b.Args) == 0 {
		return ""
	}
	return b.Args[0]
}

func (c *BuiltinCommand) GetGroup() CommandGroup {
	return GroupBuiltin
}

func (c *ExternalCommand) GetGroup() CommandGroup {
	return GroupExternal
}

func (c *BuiltinCommand) GetPath() string {
	return c.Args[0]
}

func (c *ExternalCommand) GetPath() string {
	return c.fullPath
}

func (cmd ExternalCommand) Run() {
	fmt.Printf("%s is %s\n", cmd.Args[0], cmd.fullPath)
}

type ExitCommand struct {
	BuiltinCommand
}
type EchoCommand struct {
	BuiltinCommand
}
type TypeCommand struct {
	BuiltinCommand
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
	cmd, errcmd := findCommand(c.Args[1:])
	if errcmd == nil {
		switch cmd.GetGroup() {
		case GroupBuiltin:
			fmt.Printf("%s is a shell builtin\n", c.Args[1])
		case GroupExternal:
			fmt.Printf("%s is %s\n", c.Args[1], cmd.GetPath())
		default:
		}

	} else {
		fmt.Printf("%s: not found\n", c.Args[1])
	}
}

const (
	exit  string = "exit"
	echo  string = "echo"
	type_ string = "type"
)

func BuiltinCommandsFactory(args []string) (bc IBaseCommand, err error) {
	base := BuiltinCommand{BaseCommand{Args: args}}
	switch args[0] {
	case "exit":
		return &ExitCommand{base}, nil
	case "echo":
		return &EchoCommand{base}, nil
	case "type":
		return &TypeCommand{base}, nil
	default:
		return nil, fmt.Errorf("unknown command [%s]\n", args)
	}
}

func findCommand(commandArgs []string) (command IBaseCommand, err error) {
	// look for builtin
	builtincmd, builtinerr := BuiltinCommandsFactory(commandArgs)

	if builtinerr != nil {
		// look for external commands
		path_full := os.Getenv("PATH")
		paths := strings.Split(path_full, ":")
		for _, path := range paths {
			fullpathCmd := filepath.Join(path, commandArgs[0])
			fileinfo, err := os.Stat(fullpathCmd)
			if err == nil {
				hasExecBit := fileinfo.Mode().Perm()&0111 != 0
				hasExecBit = fileinfo.Mode().IsRegular() && hasExecBit
				if !hasExecBit {
					continue
				}
				return &ExternalCommand{BaseCommand{Args: commandArgs}, fullpathCmd}, nil
			}
		}
	} else {
		return builtincmd, nil
	}
	return nil, fmt.Errorf("%s: command not found\n", commandArgs[0])
}
