package cmd

import (
	"github.com/salindersidhu/wam/curse"
	"github.com/urfave/cli"
)

var addon = curse.NewAddon()

// WamCommands returns an array of commands used by the World of Warcraft addon
// manager cli.
func WamCommands() []cli.Command {
	// Return array of CLI commands
	return []cli.Command{
		getCmd,
		initCmd,
		installCmd,
	}
}
