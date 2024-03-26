package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"dnsgo/internal/constants"
	"dnsgo/internal/remove"
)

var removeShortDisplay = constants.Strings["display_remove"]
var removeCommandName = "remove"
var removeCommandShorthand = "r"
var removeCommandAliases = []string{"rm", "clear"}

func init() {
	rootCmd.PersistentFlags().BoolP(removeCommandName, removeCommandShorthand, false, removeShortDisplay)

	rootCmd.ParseFlags(os.Args[1:])

	if cmdFlag, _ := rootCmd.PersistentFlags().GetBool(removeCommandName); cmdFlag {
		remove.Handle()
	}

	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:     removeCommandName,
	Aliases: removeCommandAliases,
	Short:   removeShortDisplay,
	Run: func(cmd *cobra.Command, args []string) {
		remove.Handle()
	},
}
