package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"dnsgo/internal/constants"
	"dnsgo/internal/current"
)

var currentShortDisplay = constants.Strings["display_current"]
var currentCommandName = "current"
var currentCommandShorthand = "c"
var currentCommandAliases = []string{"crt"}

func init() {

	rootCmd.PersistentFlags().BoolP(currentCommandName, currentCommandShorthand, false, currentShortDisplay)

	rootCmd.ParseFlags(os.Args[1:])

	if cmdFlag, _ := rootCmd.PersistentFlags().GetBool(currentCommandName); cmdFlag {
		current.Handle()
	}

	rootCmd.AddCommand(currentCmd)
}

var currentCmd = &cobra.Command{
	Use:     currentCommandName,
	Aliases: currentCommandAliases,
	Short:   currentShortDisplay,
	Run: func(cmd *cobra.Command, args []string) {
		current.Handle()
	},
}
