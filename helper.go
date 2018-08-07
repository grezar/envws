package main

import (
	"os/user"
	"strings"
)

func getAWSProfilesWithCredential(path string) (profiles []Profile, err error) {
	usr, err := user.Current()
	if err != nil {
		return profiles, err
	}
	path = strings.Replace(path, "~", usr.HomeDir, 1)
	profiles, err = parseAWSCredentials(path)
	if err != nil {
		return profiles, err
	}
	return profiles, err 
}
