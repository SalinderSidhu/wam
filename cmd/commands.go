package cmd

import (
	"github.com/urfave/cli"
)

// Commands represents a collection of CLI command objects
var Commands = []cli.Command{
	commandGet,
}

var commandGet = cli.Command{
	Name:   "get",
	Usage:  "Find and download an addon with a specific ID or name",
	Action: doGet,
}

func doGet(c *cli.Context) {}
