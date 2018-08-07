package credential

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

func LoadCredential(path string) (profiles []Profile) {
	config, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to load AWS credential file : %v", err)
		os.Exit(1)
	}

	for _, cs := range config.Sections() {
		profile := cs.Name()
		fmt.Println(profile)
		id, err := cs.GetKey("aws_access_key_id")
		if err != nil {
			fmt.Sprintf("shared credentials %s in %s did not contain aws_acces_key_id", profile, path)
		}
		fmt.Println(id)
		secret, err := cs.GetKey("aws_secret_access_key")
		fmt.Println(secret)
		if err != nil {
			fmt.Sprintf("shared credentials %s in %s did not contain aws_secret_access_key", profile, path)
		}

		// profiles = append(profiles, newProfile(profile, id.String(), secret.String()))
	}
	return profiles
}
