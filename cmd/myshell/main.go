package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		switch true {
		case strings.HasPrefix(strings.ToLower(input), "exit"):
			os.Exit(0)
		case strings.HasPrefix(strings.ToLower(input), "echo"):
			fmt.Fprint(os.Stdout, strings.Replace(input, "echo ", "", 1))
		default:
			fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
		}
	}
}
