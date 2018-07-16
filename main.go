package main

import (
	"os"

	"github.com/salindersidhu/wam/cmd"
	"github.com/urfave/cli"
)

// Variables to identify the build
var (
	Author  string
	Email   string
	Version string
)

func main() {
	app := cli.NewApp()

	// Assign cli app fields
	app.Email = Email
	app.Author = Author
	app.Version = Version
	app.Name = "wam"
	app.Usage = "Install and remove World of Warcraft addons"

	// Assign cli commands
	app.Commands = cmd.commands()

	app.Run(os.Args)
}
