package main

import (
	"os/user"
	"strings"

	"github.com/grezar/envws/credential"
)

func getAWSProfilesWithCredential(path string) (profiles []credential.Profile, err error) {
	usr, err := user.Current()
	if err != nil {
		return profiles, err
	}
	path = strings.Replace(path, "~", usr.HomeDir, 1)
	profiles = credential.LoadCredential(path)
	return profiles, err
}
