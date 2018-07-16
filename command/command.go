package cmd

import (
	"github.com/urfave/cli"
)

func Commands() []cli.Command {
	// Return array of CLI commands
	return []cli.Command{
		cli.Command{
			Name:   "init",
			Usage:  "Initialize a new addon profile",
			Action: nil,
		},
		cli.Command{
			Name:   "get",
			Usage:  "Display information about an addon",
			Action: doGet,
		},
		cli.Command{
			Name:   "install",
			Usage:  "Download and install an addon",
			Action: nil,
		},
	}
}
