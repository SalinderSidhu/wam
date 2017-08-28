package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"./addon"
	"./curse"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// CommandWrapper wraps dependencies used by CLI commands
type CommandWrapper struct {
	util     addon.Util
	strOK    string
	strINFO  string
	strERROR string
}

// Commands returns an array of CLI commands
func Commands() []cli.Command {
	// Create dependencies and Wrapper
	u := curse.NewUtil()
	w := CommandWrapper{
		util:     u,
		strOK:    color.GreenString("[OK]"),
		strINFO:  color.BlueString("[INFO]"),
		strERROR: color.RedString("[ERROR]"),
	}
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
	var notFound []string
	var addonTable [][]string
	// Create an ASCII table and set table header for addon data
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"name", "version", "updated"})
	fmt.Fprintf(color.Output, "%s Searching...\n", w.strINFO)
	for _, arg := range c.Args() {
		// Attempt to get addon data for each curse id
		data, err := w.util.GetData(arg)
		if err != nil {
			//
			notFound = append(notFound, color.MagentaString(arg))
			continue
		}
		// Append addon name, version and last updated date to the table
		addonTable = append(addonTable, []string{
			data.Name,
			data.Version,
			time.Unix(data.Epoch, 0).Format("Jan 02, 2006"),
		})
	}
	// Print the addon table data
	if len(addonTable) > 0 {
		table.AppendBulk(addonTable)
		fmt.Fprintf(color.Output, "%s Found...\n", w.strOK)
		table.Render()
	}
	// Print list of addon IDs not found
	if len(notFound) > 0 {
		fmt.Fprintf(color.Output, "%s Not Found: %s\n", w.strOK,
			strings.Join(notFound, ", "))
	}
}

func (w *CommandWrapper) doInstall(c *cli.Context) {
	for _, arg := range c.Args() {
		fmt.Fprintf(color.Output, "%s Installing %s...\n", w.strINFO,
			color.MagentaString(arg))
		// Install addon
		err := w.util.Install(arg)
		if err != nil {
			fmt.Fprintf(color.Output, "%s %s\n", w.strERROR, err.Error())
			continue
		}
		fmt.Fprintf(color.Output, "%s Installed %s\n", w.strOK,
			color.MagentaString(arg))
	}
}
