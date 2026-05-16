package main

import "fmt"

type BuiltinCommands string

const (
	exit BuiltinCommands = "exit"
)

func ParseBuiltinCommands(s string) (bc BuiltinCommands, err error) {
	cmds := map[BuiltinCommands]struct{}{
		exit: {},
	}
	cmd := BuiltinCommands(s)
	_, ok := cmds[cmd]
	if !ok {
		return bc, fmt.Errorf("cannot parse: [%s] as builtin command", bc)
	}
	return cmd, nil
}
