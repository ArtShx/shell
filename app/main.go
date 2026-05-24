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
		commandArgs := strings.Split(command, " ")
		if len(commandArgs) == 0 {
			fmt.Printf("Failed parsing command [%s]\n", command)
			continue
		}

		var redirectOutput = -1
		var redirectFile string
		var redirectStdout bool = false
		var redirectStderr bool = false

		for i, val := range commandArgs {
			if val == "1>" || val == ">" {
				redirectOutput = i
				redirectStdout = true
				break
			}
			if val == "2>" {
				redirectOutput = i
				redirectStderr = true
				break
			}
		}

		if redirectOutput != -1 {
			redirectFile = commandArgs[redirectOutput+1]
			commandArgs = commandArgs[:redirectOutput]
		}

		cmd, errcmd := findCommand(commandArgs)
		if errcmd != nil {
			fmt.Print(errcmd)
			continue
		}
		stdout, stderr, err := cmd.Run()
		if stderr != "" {
			if redirectStderr {
				err = os.WriteFile(redirectFile, []byte(stderr), 0644)
			} else {
				fmt.Fprintf(os.Stderr, "%s", stderr)
			}
		}
		if redirectStdout {
			err = os.WriteFile(redirectFile, []byte(stdout), 0644)
		} else {
			fmt.Fprintf(os.Stdout, "%s", stdout)
		}
	}
}
