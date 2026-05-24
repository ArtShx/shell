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
		var outputFile string

		for i, val := range commandArgs {
			if val == "1>" || val == ">" {
				redirectOutput = i
				break
			}
		}

		if redirectOutput != -1 {
			outputFile = commandArgs[redirectOutput+1]
			commandArgs = commandArgs[:redirectOutput]
		}

		cmd, errcmd := findCommand(commandArgs)
		if errcmd != nil {
			fmt.Print(errcmd)
			continue
		}
		stdout, stderr, err := cmd.Run()
		if stderr != "" {
			fmt.Printf("%s", stderr)
			continue
		}
		if redirectOutput == -1 {
			fmt.Printf("%s", stdout)
		} else {
			err = os.WriteFile(outputFile, []byte(stdout), 0644)
			_ = err
		}
	}
}
