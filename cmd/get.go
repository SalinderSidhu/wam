package cmd

import (
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/salindersidhu/wam/util"
	"github.com/urfave/cli"
)

var getCmd = cli.Command{
	Name:   "get",
	Usage:  "Display information about an addon",
	Action: doGet,
}

func doGet(ctx *cli.Context) {
	var notFound []string
	var addonTable [][]string
	// Create an ASCII table to show addon data
	table := tablewriter.NewWriter(os.Stdout)
	if ctx.NArg() > 0 {
		// Set ASCII table header for addon data
		table.SetHeader([]string{"name", "version", "updated"})
		util.PrintfOk("Searching...\n")
	}
	for _, arg := range ctx.Args() {
		// Attempt to get addon data for each curse id
		data, err := addon.InitMetadata(arg)
		if err != nil {
			// If an error occurred, add id to not found list
			notFound = append(notFound, color.MagentaString(arg))
			continue
		}
		// Append addon name, version and last updated date to the table
		addonTable = append(addonTable, []string{
			data.Name,
			data.Version,
			time.Unix(data.Date, 0).Format("Jan 02, 2006"),
		})
	}
	// Print the addon table data
	if len(addonTable) > 0 {
		table.AppendBulk(addonTable)
		util.PrintfOk("Found...\n")
		table.Render()
	}
	// Print list of addon IDs not found
	if len(notFound) > 0 {
		util.PrintfOk("Not Found: %s\n", strings.Join(notFound, ", "))
	}
}
