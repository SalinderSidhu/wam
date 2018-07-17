package cmd

import (
	"github.com/fatih/color"
	"github.com/salindersidhu/wam/util"
	"github.com/urfave/cli"
)

var installCmd = cli.Command{
	Name:   "install",
	Usage:  "Download and install an addon",
	Action: doInstall,
}

func doInstall(ctx *cli.Context) {
	for _, arg := range ctx.Args() {
		util.PrintfInfo("Installing %s...\n", color.MagentaString(arg))
		// Install addon
		if err := curse.Install(arg); err != nil {
			util.PrintfError("%s\n", err.Error())
			continue
		}
		util.PrintfOk("Installed %s\n", color.MagentaString(arg))
	}
}
