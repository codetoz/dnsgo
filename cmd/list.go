package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"dnsgo/internal/list"
	"dnsgo/internal/references"
)

func init() {
	rootCmd.PersistentFlags().BoolP("list", "l", false, references.Strings["display_list"])

	rootCmd.ParseFlags(os.Args[1:])

	if cmdFlag, _ := rootCmd.PersistentFlags().GetBool("list"); cmdFlag {
		list.Handle()
	}

	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   references.Strings["display_list"],
	Run: func(cmd *cobra.Command, args []string) {
		list.Handle()
	},
}
