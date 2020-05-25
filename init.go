package main

import (
	"bufio"
	"os"
)

func Initialize() {
	_ = os.Chdir(os.Getenv("HOMEPATH"))
	cwd := os.Getenv("HOMEPATH")
	_ = os.Setenv("TURT_CURDIR", cwd)
	_ = os.Setenv("TURT_PASTDIR", cwd)

	readResources()
}

func readResources() {
	if rscFile, err := os.Open(os.Getenv("HOMEPATH") + ".turtlersc"); err == nil {
		rscReader := bufio.NewReader(rscFile)
		for {
			line, _ := rscReader.ReadString('\n')
			_ = Execute(line)

			if _, complete := rscReader.Peek(1); complete != nil {
				break
			}
		}
	}
}
