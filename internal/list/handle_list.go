package list

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

func Handle() {
	green := color.New(color.FgGreen).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)

	fmt.Fprintf(w, "%s\t%s\t%s\n", green("Name"), white("IPs"), red("Address"))

	for _, option := range DnsOptions {
		fmt.Fprintf(w, "%s\t%s\t%s\n", green(option.Name), white(option.IPs), red(option.Address))
	}

	w.Flush()

	os.Exit(0)
}
