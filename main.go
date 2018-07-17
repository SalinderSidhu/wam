package main

import (
	"log"
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

	// Assign wam commands to cli app
	app.Commands = cmd.WamCommands()

	// Run the application with user arguments
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
