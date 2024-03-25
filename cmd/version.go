package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"dnsgo/internal/references"
	"dnsgo/internal/version"
)

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, references.Strings["display_version"])

	rootCmd.ParseFlags(os.Args[1:])

	if versionFlag, _ := rootCmd.PersistentFlags().GetBool("version"); versionFlag {
		version.HandleVersion()
	}

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: references.Strings["display_version"],
	Run: func(cmd *cobra.Command, args []string) {
		version.HandleVersion()
	},
}
