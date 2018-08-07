package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var profileNameRegexp = regexp.MustCompile("^\\[([A-Za-z0-9_-]+)\\]$")

func parseAWSCredentials(path string) (profiles []Profile, err error) {
	file, err := os.Open(path)
	if err != nil {
		return profiles, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	var profile Profile
	for s.Scan() {
		text := s.Text()
		fields := strings.Fields(text)
		if profileNameRegexp.MatchString(text) {
			matches := profileNameRegexp.FindAllStringSubmatch(text, -1)
			if err != nil {
				return profiles, err
			}
			for _, match := range matches {
				profile = Profile{}
				profile.Name = match[1]
			}
		} else if strings.Contains(text, "aws_access_key_id") {
			profile.Credential.AWSAccessKeyId = fields[len(fields)-1]
		} else if strings.Contains(text, "aws_secret_access_key") {
			profile.Credential.AWSSecretAccessKey = fields[len(fields)-1]
			profiles = append(profiles, profile)
		} else {
			continue
		}
	}
	return profiles, err
}
