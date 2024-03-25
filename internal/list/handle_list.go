package list

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func Handle() {
	green := color.New(color.FgGreen).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)

	termWidth, _, _ := term.GetSize(0)

	if termWidth > 80 {
		fmt.Fprintf(w, "%s\t%s\t%s\n", green("Name"), white("IPs"), red("Address"))
		for _, option := range DnsOptions {
			fmt.Fprintf(w, "%s\t%s\t%s\n", green(option.Name), white(option.IPs), red(option.Address))
		}
	} else {
		fmt.Fprintf(w, "%s\t\t\t%s\n", green("Name"), white("IPs"))
		for _, option := range DnsOptions {
			fmt.Fprintf(w, "%s\t\t\t%s\n", green(option.Name), white(option.IPs))
		}
	}

	w.Flush()

	os.Exit(0)
}
