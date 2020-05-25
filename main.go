// Turtle - A shell with something alive inside

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	history []string
)

func main() {
	Initialize()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		if _ = Execute(input); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}

func Execute(input string) error {
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	history = append(history, input)

	for keyword, macro := range Aliased {
		input = strings.ReplaceAll(input, keyword, macro)
	}

	args := strings.Split(input, " ")

	if strings.Index(args[0], "!") != 0 {
		args = Turtle(args)
	} else {
		args[0] = args[0][1:]
	}

	switch args[0] {
	case "alias":
		if args[2] != "=" {
			return errors.New("syntax [keyword] = [macro string]")
		}

		Alias(args[1], strings.Join(args[3:], " "))
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}

		return ChangeDir(args[1])
	case "exit":
		os.Exit(0)
	case "history":
		for _, entry := range history {
			fmt.Println(entry)
		}
	case "ls":
		ListDir()
		return nil
	case "pwd":
		fmt.Println(WorkingDir())
		return nil
	case "unalias":
		Unalias(args[1])
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
