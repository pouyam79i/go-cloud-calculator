package cmd

import (
	"fmt"
	"go-cloud-calculator/internal/pkg/echoz"
	"strings"

	"github.com/spf13/cobra"
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echo is used to build web server",
	Long: `go-cloud-calculator echo 'port'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			fmt.Println("Launching echo server...")
			var port string = strings.TrimSpace(args[0])
			echoz.BuildServer(port)
		} else {
			fmt.Println("Wrong number of input args")
		}
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
