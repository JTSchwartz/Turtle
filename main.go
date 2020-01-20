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

func main() {
	Initialize()

	reader := bufio.NewReader(os.Stdin)
	for {
		// fmt.Print("𝛙 ")
		fmt.Print("🐢 ")
		input, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		if _ = execute(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execute(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}

		return ChangeDir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
