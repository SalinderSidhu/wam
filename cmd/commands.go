package cmd

import (
	"fmt"

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
		fmt.Printf("[INFO] searching for %s...\n", arg)
		// Get data for each addon
		addon, err := w.util.GetData(arg)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return
		}
		// Determine action (display info or download) based on flag
		if c.Bool("info") {
			// Print the addon's name and version
			fmt.Printf("[OK] found %s %s\n", addon.Name, addon.Version)
		} else {
			// Download the addon
			_, err := w.util.Download(addon)
			if err != nil {
				fmt.Printf("[ERROR] %s\n", err.Error())
				return
			}
			// Print the addon's name and version
			fmt.Printf("[OK] downloaded %s %s\n", addon.Name, addon.Version)
		}
	}
}
