package cmd

import (
	"github.com/moleske/picture-shuffler/shuffler"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "picture-shuffler",
	Short: "Display shuffled pictures",
	Long: `Takes a directory and walks the entire structure.
Then tries to display every file found`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.ExactArgs(1),
	Run:  shuffler.Shuffle,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.picture-shuffler.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
//func initConfig() {
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		home, err := os.UserHomeDir()
//		cobra.CheckErr(err)
//
//		// Search config in home directory with name ".picture-shuffler" (without extension).
//		viper.AddConfigPath(home)
//		viper.SetConfigType("yaml")
//		viper.SetConfigName(".picture-shuffler")
//	}
//
//	viper.AutomaticEnv() // read in environment variables that match
//
//	// If a config file is found, read it in.
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
//	}
//}
