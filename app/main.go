package main

import (
	"bufio"
	"fmt"
	"os"
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

		builtincmd, builtinerr := ParseBuiltinCommands(command)

		if builtinerr != nil {
			// try not builtin cmd

		} else {
			switch builtincmd {
			case "exit":
				return
			}
		}

		fmt.Printf("%s: command not found\r\n", command)
	}
}
