package cmd

import (
	"github.com/fatih/color"
	"github.com/salindersidhu/wam/util"
	"github.com/urfave/cli"
)

var initCmd = cli.Command{
	Name:   "init",
	Usage:  "Initialize a new addon profile",
	Action: doInit,
}

func doInit(ctx *cli.Context) {
	util.PrintfOk("Creating profile...\n")
	// Obtain path from argument (if any) and create the profile
	if err := curse.InitProfile(ctx.Args().First()); err != nil {
		util.PrintfError("%s\n", err.Error())
		return
	}
	util.PrintfOk("Profile created in %s\n", color.YellowString("wam.json"))
}
