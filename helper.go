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

func judgeWetherEvalCommandOrNot(cmd string) bool {
    for _, ec := range evalCommands {
        if ec == cmd {
            return true
        }
    }
    return false
}
