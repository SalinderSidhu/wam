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
			Usage:  "Display information about an addon",
			Action: w.doGet,
		},
		cli.Command{
			Name:   "install",
			Usage:  "Download and install an addon",
			Action: w.doInstall,
		},
	}
}

func (w *CommandWrapper) doGet(c *cli.Context) {
	for _, arg := range c.Args() {
		fmt.Printf("[INFO] searching for %s...\n", arg)
		// Get addon information
		info, err := w.util.GetInfo(arg)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			continue
		}
		fmt.Printf("[OK] found %s\n", info)
	}
}

func (w *CommandWrapper) doInstall(c *cli.Context) {
	for _, arg := range c.Args() {
		fmt.Printf("[INFO] searching for %s...\n", arg)
		// Download the addon
		err := w.util.Download(arg)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			continue
		}
		fmt.Printf("[OK] downloaded %s\n", arg)
	}
}
