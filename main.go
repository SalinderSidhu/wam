package main

import (
	"os"

	"./cmd"
	"github.com/urfave/cli"
)

// Variables to identify the build
var (
	Name    string
	Version string
)

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Salinder Sidhu"
	app.Email = "salinder.sid@gmail.com"
	app.Usage = "Manage World of Warcraft addons"
	app.Commands = cmd.Commands()

	app.Run(os.Args)
}
