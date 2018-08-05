package main

type Profile struct {
	Name       string
	Credential credential
}

type credential struct {
	AWSAccessKeyId     string
	AWSSecretAccessKey string
}
