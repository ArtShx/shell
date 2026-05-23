package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type IBaseCommand interface {
	Run() string
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

func (cmd ExternalCommand) Run() string {
	execCmd := exec.Command(cmd.Args[0], cmd.Args[1:]...)
	out, _ := execCmd.CombinedOutput()
	return fmt.Sprintf("%s", out)
	// if err != nil {
	// 	fmt.Println(err)
	// }
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
type PwdCommand struct {
	BuiltinCommand
}
type CdCommand struct {
	BuiltinCommand
}

func (c ExitCommand) Run() string {
	os.Exit(0)
	return ""
}

func (c EchoCommand) Run() string {
	return fmt.Sprintln(strings.Join(c.Args[1:], " "))
}

func (c TypeCommand) Run() string {
	if len(c.Args) <= 1 {
		return ""
	}
	cmd, errcmd := findCommand(c.Args[1:])
	if errcmd == nil {
		switch cmd.GetGroup() {
		case GroupBuiltin:
			return fmt.Sprintf("%s is a shell builtin\n", c.Args[1])
		case GroupExternal:
			return fmt.Sprintf("%s is %s\n", c.Args[1], cmd.GetPath())
		default:
		}

	}
	return fmt.Sprintf("%s: not found\n", c.Args[1])
}

func (c PwdCommand) Run() string {
	nav := GetNavigation()
	return fmt.Sprintf("%s\n", nav.wd)
}

func (c CdCommand) Run() string {
	if len(c.Args) > 2 {
		return ""
	}
	ChangeDirectory(c.Args[1])
	return ""
}

const (
	exit  string = "exit"
	echo  string = "echo"
	type_ string = "type"
	pwd   string = "pwd"
	cd    string = "cd"
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
	case "pwd":
		return &PwdCommand{base}, nil
	case "cd":
		return &CdCommand{base}, nil
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
