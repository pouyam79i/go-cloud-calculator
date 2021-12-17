package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// Help message
const HelpMessage = `Go Cloud Calculator!
  This command is used to build a server or client!
  Server is able to repond to client.
  Client is able to send requests to the client!

  If you want to create server, take following command as an example:
  go-cloud-calculator server 8000
  - note: 8000 is an available port.

  If you want to create client, take following command as an example:
  go-cloud-calculator client 127.0.0.1:8000
  - note: 127.0.0.1 is the ip address and 8000 is the port you want to connect!

  remember to launch server before client.
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cloud-calculator",
	Short: "Go Cloud Calculator",
	Long:  HelpMessage,
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Wellcome to Go Cloud Calculator")
		helpFlag, _ := cmd.Flags().GetBool("help")
		if helpFlag {
			fmt.Printf("%s \n", HelpMessage)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cloud-calculator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("help", "h", false, "Help message")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-cloud-calculator" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-cloud-calculator")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
