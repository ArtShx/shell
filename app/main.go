package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Failed reading input")
		return
	}
	command = strings.TrimSuffix(command, "\r\n")
	fmt.Printf("%s: command not found\r\n", command)
}
