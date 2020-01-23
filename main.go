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

		if _ = Execute(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func Execute(input string) error {
	input = strings.TrimSuffix(input, "\n")

	aliasedKeywords := make([]string, 0, len(Aliased))
	for _, key := range Aliased {
		aliasedKeywords = append(aliasedKeywords, key)
	}

	for _, keyword := range aliasedKeywords {
		input = strings.ReplaceAll(input, keyword, Aliased[keyword])
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
	case "unalias":
		Unalias(args[1])
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
