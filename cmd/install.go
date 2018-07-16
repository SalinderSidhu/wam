package cmd

import (
	"github.com/salindersidhu/wam/util"
	"github.com/urfave/cli"
)

var installCmd = cli.Command{
	Name:   "install",
	Usage:  "Download and install an addon",
	Action: doInstall,
}

func doInstall(ctx *cli.Context) {
	util.PrintfError("Action not implemented\n")
}
