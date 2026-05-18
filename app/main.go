package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Failed reading input")
			return
		}
		command = strings.TrimSuffix(command, "\r\n") // windows
		command = strings.TrimSuffix(command, "\n")
		command_args := strings.Split(command, " ")
		if len(command_args) == 0 {
			fmt.Printf("Failed parsing command [%s]\n", command)
			continue
		}

		builtincmd, builtinerr := ParseBuiltinCommands(command_args)
		found_command := false

		if builtinerr != nil {
			// try not builtin cmd
			path_full := os.Getenv("PATH")
			paths := strings.Split(path_full, ":")
			for _, path := range paths {
				fullpathCmd := filepath.Join(path, command_args[0])
				fileinfo, err := os.Stat(fullpathCmd)
				if err == nil {
					hasExecBit := fileinfo.Mode().Perm()&0111 != 0
					hasExecBit = fileinfo.Mode().IsRegular() && hasExecBit
					fmt.Printf("%s is %s\n", command_args[0], fullpathCmd)
					found_command = true
					break
				}
			}
		} else {
			builtincmd.Run()
			continue
		}

		if !found_command {
			fmt.Printf("%s: command not found\r\n", command)
		}
	}
}
