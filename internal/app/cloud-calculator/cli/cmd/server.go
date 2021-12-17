package cmd

import (
	"fmt"
	"strings"

	"go-cloud-calculator/internal/pkg/server"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server command",
	Long:  `go-cloud-calculator server 'port'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			fmt.Println("Launching server...")
			var port string = strings.TrimSpace(args[0])
			server.RunServer(port)
		} else {
			fmt.Println("Wrong number of input args")
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
