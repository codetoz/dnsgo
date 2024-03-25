package list

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func Handle() {
	white := color.New(color.FgWhite).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)

	termWidth, _, _ := term.GetSize(0)

	twoColPattern := "| %s\t\t\t | %s\n"
	threeColPattern := "| %s\t | %s\t | %s\n"

	if termWidth > 80 {
		fmt.Fprintf(w, threeColPattern,
			white("Name"),
			white("IPs"),
			white("Address"))

		for idx, option := range DnsOptions {
			if idx%2 == 0 {
				fmt.Fprintf(w, threeColPattern,
					cyan(option.Name),
					cyan(option.IPs),
					cyan(option.Address))
			} else {
				fmt.Fprintf(w, threeColPattern,
					green(option.Name),
					green(option.IPs),
					green(option.Address),
				)
			}
		}
	} else {
		fmt.Fprintf(w, twoColPattern,
			white("Name"),
			white("IPs"))

		for idx, option := range DnsOptions {
			if idx%2 == 0 {
				fmt.Fprintf(w, twoColPattern,
					cyan(option.Name),
					cyan(option.IPs),
				)
			} else {
				fmt.Fprintf(w, twoColPattern,
					green(option.Name),
					green(option.IPs),
				)
			}
		}
	}

	w.Flush()

	os.Exit(0)
}
