package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"dnsgo/internal/references"
	"dnsgo/internal/version"
)

var versionShortDisplay = references.Strings["display_version"]
var versionCommandName = "version"
var versionCommandShorthand = "v"

// var versionCommandAliases = []string{"vv"}

func init() {
	rootCmd.PersistentFlags().BoolP(versionCommandName, versionCommandShorthand, false, versionShortDisplay)

	rootCmd.ParseFlags(os.Args[1:])

	if cmdFlag, _ := rootCmd.PersistentFlags().GetBool(versionCommandName); cmdFlag {
		version.Handle()
	}

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   versionCommandName,
	Short: versionShortDisplay,
	Run: func(cmd *cobra.Command, args []string) {
		version.Handle()
	},
}
