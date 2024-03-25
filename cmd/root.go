package cmd

import (
	"dnsgo/internal/version"

	"github.com/spf13/cobra"
)

// // Used for flags.
// var cfgFile string

// var userLicense string

var rootCmd = &cobra.Command{
	Use:   "dnsgo",
	Short: "DNS Go",
	Long:  "DNS Go version",
}

// Execute executes the root command.
func Execute() error {
	if versionFlag, _ := rootCmd.PersistentFlags().GetBool("version"); versionFlag {
		version.HandleVersion()
	}

	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().String("version", "v", "v")

	// rootCmd.PersistentFlags().StringP("author", "a", "Mahdi Darabi", "author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "Mahdi Darabi mahdidarabi7@yahoo.com")
	// viper.SetDefault("license", "apache")

	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)
}

func initConfig() {
	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	home, err := os.UserHomeDir()
	// 	cobra.CheckErr(err)

	// 	// Search config in home directory with name ".cobra" (without extension).
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigType("yaml")
	// 	viper.SetConfigName(".cobra")
	// }

	// viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// }
}
