package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"dnsgo/internal/version"
)

func init() {

	// Add a flag for version
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Display the version number")

	rootCmd.ParseFlags(os.Args[1:])

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version number",
	Long:  `version number og dnsgo`,
	Run: func(cmd *cobra.Command, args []string) {
		version.HandleVersion()
	},
}
