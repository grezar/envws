package main

import (
	"os/user"
	"strings"
)

func getUserHomeDir(path string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	path = strings.Replace(path, "~", usr.HomeDir, 1)
	return path
}
