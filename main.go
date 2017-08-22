package main

import (
	"os"

	"./cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = cmd.Name
	app.Email = cmd.Email
	app.Author = cmd.Author
	app.Version = cmd.Version
	app.Commands = cmd.Commands()

	app.Run(os.Args)
}
