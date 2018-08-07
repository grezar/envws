package credential

type Profile struct {
	Name       string
	Credential credential
}

type credential struct {
	AWSAccessKeyId     string
	AWSSecretAccessKey string
}

func newProfile(profile, id, secret string) Profile {
	return Profile{
		Name: profile,
		Credential: credential{
			AWSAccessKeyId:     id,
			AWSSecretAccessKey: secret,
		},
	}
}
