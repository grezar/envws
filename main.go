package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "envws"
	app.Usage = "Handle AWS credentials with environment variables"
	app.Version = "0.1.0"

	app.Run(os.Args)
}
