package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"dnsgo/internal/constants"
	"dnsgo/internal/list"
	"dnsgo/internal/set"
)

var setShortDisplay = constants.Strings["display_set"]
var setCommandName = "set"
var setCommandShorthand = "s"

func init() {
	rootCmd.PersistentFlags().StringP(setCommandName, setCommandShorthand, "", setShortDisplay)

	rootCmd.ParseFlags(os.Args[1:])

	if cmdFlag, _ := rootCmd.PersistentFlags().GetString(setCommandName); cmdFlag != "" {
		flagValue, _ := rootCmd.Flags().GetString(setCommandName)
		set.Handle(flagValue)
	}

	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   setCommandName + " <dns-name>",
	Short: setShortDisplay,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			set.Handle(args[0])
		} else {
			fmt.Println(constants.Strings["enter_dns_name"])
			list.Handle()
		}
	},
}
