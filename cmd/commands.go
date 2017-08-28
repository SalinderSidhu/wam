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
	var notFound []string
	var addonTable [][]string
	// Create an ASCII table and set table header for addon data
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"name", "version", "updated"})
	w.fprintcInfo("Searching...\n")
	for _, arg := range c.Args() {
		// Attempt to get addon data for each curse id
		data, err := w.util.GetData(arg)
		if err != nil {
			// If an error occured, add id to not found list
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
		w.fprintcOk("Found...\n")
		table.Render()
	}
	// Print list of addon IDs not found
	if len(notFound) > 0 {
		w.fprintcOk("Not Found: %s\n", strings.Join(notFound, ", "))
	}
}

func (w *CommandWrapper) doInstall(c *cli.Context) {
	for _, arg := range c.Args() {
		w.fprintcInfo("Installing %s...\n", color.MagentaString(arg))
		// Install addon
		err := w.util.Install(arg)
		if err != nil {
			w.fprintcError("%s\n", err.Error())
			continue
		}
		w.fprintcOk("Installed %s\n", color.MagentaString(arg))
	}
}

func (w *CommandWrapper) fprintcOk(f string, a ...interface{}) (int, error) {
	return w.fprintc("%s %s", color.GreenString("[OK]"), fmt.Sprintf(f, a...))
}

func (w *CommandWrapper) fprintcInfo(f string, a ...interface{}) (int, error) {
	return w.fprintc("%s %s", color.BlueString("[INFO]"), fmt.Sprintf(f, a...))
}

func (w *CommandWrapper) fprintcError(f string, a ...interface{}) (int, error) {
	return w.fprintc("%s %s", color.RedString("[ERROR]"), fmt.Sprintf(f, a...))
}

func (w *CommandWrapper) fprintc(f string, a ...interface{}) (int, error) {
	return fmt.Fprintf(color.Output, f, a...)
}
