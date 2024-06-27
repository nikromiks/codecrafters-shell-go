package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands = map[string]string{
	"exit": "exit",
	"echo": "echo",
	"type": "type",
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		switch true {
		case strings.HasPrefix(strings.ToLower(input), "exit"):
			os.Exit(0)

		case strings.HasPrefix(strings.ToLower(input), "echo"):
			fmt.Fprint(os.Stdout, strings.TrimPrefix(input, "echo "))

		case strings.HasPrefix(strings.ToLower(input), "type"):
			rest := strings.TrimPrefix(input, "type")
			rest = strings.Trim(rest, "\n ")
			if _, ok := commands[rest]; !ok {
				fmt.Fprintln(os.Stdout, rest+": not found")
				continue
			}
			fmt.Fprintln(os.Stdout, rest+" is a shell builtin")

		default:
			fmt.Fprintln(os.Stdout, input[:len(input)-1]+": command not found")
		}
	}
}
