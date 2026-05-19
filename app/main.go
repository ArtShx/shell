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
		command_args := strings.Split(command, " ")
		if len(command_args) == 0 {
			fmt.Printf("Failed parsing command [%s]\n", command)
			continue
		}

		cmd, errcmd := findCommand(command_args)
		if errcmd != nil {
			fmt.Print(errcmd)
			continue
		}
		cmd.Run()
	}
}
