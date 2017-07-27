package main

import (
	"os"

	"./cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wam"
	app.Author = "Salinder"
	app.Email = "salinder.sid@gmail.com"
	app.Version = cmd.Version
	app.Commands = cmd.Commands()

	app.Run(os.Args)
}
