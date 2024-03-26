package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"dnsgo/internal/current"
	"dnsgo/internal/references"
)

func init() {
	rootCmd.PersistentFlags().BoolP("current", "c", false, references.Strings[""])

	rootCmd.ParseFlags(os.Args[1:])

	if cmdFlag, _ := rootCmd.PersistentFlags().GetBool("current"); cmdFlag {
		current.Handle()
	}

	rootCmd.AddCommand(currentCmd)
}

var currentCmd = &cobra.Command{
	Use:     "current",
	Aliases: []string{"crt"},
	Short:   references.Strings[""],
	Run: func(cmd *cobra.Command, args []string) {
		current.Handle()
	},
}
