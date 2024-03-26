package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"dnsgo/internal/list"
	"dnsgo/internal/references"
)

var listShortDisplay = references.Strings["display_list"]
var listCommandName = "list"
var listCommandShorthand = "l"
var listCommandAliases = []string{"ls"}

func init() {
	rootCmd.PersistentFlags().BoolP(listCommandName, listCommandShorthand, false, listShortDisplay)

	rootCmd.ParseFlags(os.Args[1:])

	if cmdFlag, _ := rootCmd.PersistentFlags().GetBool(listCommandName); cmdFlag {
		list.Handle()
	}

	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     listCommandName,
	Aliases: listCommandAliases,
	Short:   listShortDisplay,
	Run: func(cmd *cobra.Command, args []string) {
		list.Handle()
	},
}
