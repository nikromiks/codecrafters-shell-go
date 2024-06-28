package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type (
	cmdFnc func([]string)
)

var commands = make(map[string]cmdFnc)

func main() {
	initCommands()
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		in, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("error reading from stdin: %s", err.Error())
			os.Exit(1)
		}
		inputs := strings.Split(strings.TrimSpace(in), " ")

		cmd, args, builtin := getCommand(inputs)
		if builtin {
			cmdFn := commands[cmd]
			cmdFn(args)

			continue
		}
		if cmd != "" {
			command := exec.Command(cmd, args...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			err := command.Run()
			if err != nil {
				fmt.Println(err)
			}

			continue
		}
		notFound(inputs[0])
	}
}

func registerCommand(cmd string, fn cmdFnc) {
	commands[cmd] = fn
}

func initCommands() {
	registerCommand("exit", exit)
	registerCommand("echo", echo)
	registerCommand("type", typer)
}

func notFound(cmd string) {
	fmt.Printf("%s: command not found\n", cmd)
}
func exit(args []string) {
	if len(args) == 0 {
		os.Exit(1)
	}
	if code, err := strconv.Atoi(args[0]); err == nil {
		os.Exit(code)
	}
}
func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func typer(args []string) {
	cmd, _, builtin := getCommand(args)
	if builtin {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}
	if cmd != "" {
		fmt.Println(cmd)
		return
	}
	if len(args) != 0 {
		fmt.Printf("%s: not found\n", args[0])
		return
	}
	fmt.Printf("%s: not found\n", cmd)
}

func getCommand(args []string) (string, []string, bool) {
	if len(args) == 0 {
		return "", []string{}, false
	}

	cmd := args[0]

	_, builtin := commands[cmd]
	if builtin {
		return cmd, args[1:], true
	}

	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		fp := filepath.Join(path, cmd)
		if _, err := os.Stat(fp); err == nil {
			return fp, args[1:], false
		}
	}

	return "", []string{}, false
}
