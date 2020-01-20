package main

import (
	"os"
)

func Initialize() {
	os.Chdir(os.Getenv("HOMEPATH"))
	cwd, _ := os.Getwd()
	os.Setenv("TURT_CWD", cwd)
	os.Setenv("TURT_PWD", cwd)
}
