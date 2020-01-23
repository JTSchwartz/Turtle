package main

import (
	"os"
	"strings"
)

var (
	Aliased = make(map[string]string)
)

func Alias(keyword string, macro string) {
	Aliased[keyword] = macro
}

func ChangeDir(location string) error {
	switch location {
	case "~":
		location = os.Getenv("HOMEPATH")
	case "..":
		location = os.Getenv("TURT_CWD")[0:strings.LastIndex(os.Getenv("TURT_CWD"), "/")]
	case "@":
		location = os.Getenv("TURT_PWD")
	default:
		location = strings.ReplaceAll(location, "\\", "/")
	}

	os.Setenv("TURT_PWD", os.Getenv("TURT_CWD"))
	os.Setenv("TURT_CWD", location)
	return os.Chdir(location)
}

func Unalias(keyword string) {
	delete(Aliased, keyword)
}
