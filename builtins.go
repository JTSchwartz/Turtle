package main

import (
	"os"
	"strings"
)

func ChangeDir(location string) error {
	switch location {
	case "~":
		return os.Chdir(os.Getenv("HOMEPATH"))
	case "..":
		parent := strings.LastIndex(location, "/")
		return os.Chdir(location[0:parent])
	}

	location = strings.ReplaceAll(location, "\\", "/")
	return os.Chdir(location)
}
