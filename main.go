package main

import (
	"os"

	"github.com/salindersidhu/wam/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wam"
	app.Version = "0.0.9"
	app.Author = "Salinder Sidhu"
	app.Email = "salinder.sid@gmail.com"
	app.Usage = "Manage World of Warcraft addons"
	app.Commands = cmd.Commands()

	app.Run(os.Args)
}
