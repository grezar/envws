package main

import (
	"fmt"
	"os"
    "strings"

	"github.com/go-ini/ini"
)

const (
	iniDefaultHeader       = "DEFAULT"
    profilePrefixWithSpace = "profile "
)

func loadCredentialFile(path string) (config *ini.File) {
	config, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to load AWS credential file : %v", err)
		os.Exit(1)
	}
	return config
}

func loadProfile(name, path string) Profile {
	config := loadCredentialFile(path)
	profile, err := config.GetSection(name)
	if err != nil {
		fmt.Sprintf("failed to load profile %s in %s", profile, path)
		os.Exit(1)
	}
	id, err := profile.GetKey("aws_access_key_id")
	if err != nil {
		fmt.Sprintf("shared credentials %s in %s did not contain aws_acces_key_id", profile, path)
		os.Exit(1)
	}
	secret, err := profile.GetKey("aws_secret_access_key")
	if err != nil {
		fmt.Sprintf("shared credentials %s in %s did not contain aws_secret_access_key", profile, path)
		os.Exit(1)
	}
	return newProfile(profile.Name(), id.String(), secret.String())
}

func loadProfileNames(path string) (profileNames []string) {
	config := loadCredentialFile(path)
	for _, cs := range config.Sections() {
		profileName := cs.Name()
        if strings.HasPrefix(profileName, profilePrefixWithSpace) {
          continue
        }
		if profileName != iniDefaultHeader {
			profileNames = append(profileNames, profileName)
		}
	}
	return profileNames
}
