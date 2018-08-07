package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const (
	name        = "envws"
	description = "Handle AWS credentials with environment varialbes"
	version     = "0.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Usage = description
	app.Version = version

	app.Commands = []cli.Command{
		{
			Name:    "set",
			Aliases: []string{"s"},
			Usage:   "Set selected profile credentials as environment varialbes",
			Action:  cmdSet,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "profile, p",
					Value: "default",
					Usage: "AWS profile",
				},
				cli.StringFlag{
					Name:  "path",
					Value: "~/.aws/credentials",
					Usage: "AWS credential file path",
				},
			},
		},
		{
			Name:    "unset",
			Aliases: []string{"u"},
			Usage:   "Unset AWS credential environment varialbes",
			Action:  cmdUnset,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List aws profiles",
			Action:  cmdList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path",
					Value: "~/.aws/credentials",
					Usage: "AWS credential file path",
				},
			},
		},
	}
	app.Run(os.Args)
}

func cmdSet(c *cli.Context) error {
	profiles, err := getAWSProfilesWithCredential(c.String("path"))
	if err != nil {
		return err
	}
	for _, p := range profiles {
		if p.Name == c.String("profile") {
			fmt.Printf("export AWS_ACCESS_KEY_ID=%s ;", string(p.Credential.AWSAccessKeyId))
			fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s", string(p.Credential.AWSSecretAccessKey))
		}
	}
	return nil
}

func cmdUnset(c *cli.Context) {
	fmt.Println("unset AWS_ACCESS_KEY_ID ; unset AWS_SECRET_ACCESS_KEY")
}

func cmdList(c *cli.Context) error {
	profiles, err := getAWSProfilesWithCredential(c.String("path"))
	if err != nil {
		return err
	}
	for _, p := range profiles {
		fmt.Println(p.Name)
	}
	return nil
}
