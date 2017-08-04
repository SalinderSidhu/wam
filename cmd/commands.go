package cmd

import (
	"fmt"
	"log"

	"./addon"
	"./curse"
	"github.com/urfave/cli"
)

// CommandWrapper wraps dependencies used by CLI commands
type CommandWrapper struct {
	util addon.Util
}

// Commands returns an array of CLI commands
func Commands() []cli.Command {
	// Create dependencies and Wrapper
	u := curse.NewUtil()
	w := CommandWrapper{util: u}
	// Return array of CLI commands
	return []cli.Command{
		cli.Command{
			Name:   "get",
			Usage:  "Download an addon with a specific ID or name",
			Action: w.doGet,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "info, i",
					Usage: "Display addon information",
				},
			},
		},
	}
}

func (w *CommandWrapper) doGet(c *cli.Context) {
	for _, arg := range c.Args() {
		// Get data for each addon
		a, err := w.util.GetData(arg)
		if err != nil {
			log.Fatal(err)
		}
		// Print the name and version
		fmt.Printf("Addon: %s %s\n", a.Name, a.Version)
	}
}
