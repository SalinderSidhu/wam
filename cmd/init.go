package cmd

import (
	"github.com/salindersidhu/wam/util"
	"github.com/urfave/cli"
)

var initCmd = cli.Command{
	Name:   "init",
	Usage:  "Initialize a new addon profile",
	Action: doInit,
}

func doInit(ctx *cli.Context) {
	util.PrintfError("Action not implemented\n")
}
