package cmd

import (
	"fmt"
	"time"

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
		// Get addon data
		data, err := w.util.GetData(arg)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			continue
		}
		// Print addon information to std out
		fmt.Printf("[OK] found %s (%s) Updated: %s\n", data.Name, data.Version,
			time.Unix(data.Epoch, 0).Format(time.RFC822Z))
	}
}

func (w *CommandWrapper) doInstall(c *cli.Context) {
	for _, arg := range c.Args() {
		fmt.Printf("[INFO] installing %s...\n", arg)
		// Install addon
		err := w.util.Install(arg)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			continue
		}
		fmt.Printf("[OK] installed %s\n", arg)
	}
}
