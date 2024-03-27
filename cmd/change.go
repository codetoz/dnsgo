package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"dnsgo/internal/change"
	"dnsgo/internal/constants"
)

var changeShortDisplay = constants.Strings["display_change"]
var changeCommandName = "change"
var changeCommandShorthand = "C"
var changeCommandAliases = []string{"alter", "alt"}

func init() {
	rootCmd.PersistentFlags().BoolP(changeCommandName, changeCommandShorthand, false, changeShortDisplay)

	rootCmd.ParseFlags(os.Args[1:])

	if cmdFlag, _ := rootCmd.PersistentFlags().GetBool(changeCommandName); cmdFlag {
		change.Handle()
	}

	rootCmd.AddCommand(changeCmd)
}

var changeCmd = &cobra.Command{
	Use:     changeCommandName,
	Short:   changeShortDisplay,
	Aliases: changeCommandAliases,
	Run: func(cmd *cobra.Command, args []string) {
		change.Handle()
	},
}
