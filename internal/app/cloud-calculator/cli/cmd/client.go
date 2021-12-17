package cmd

import (
	"fmt"
	"strings"

	"go-cloud-calculator/internal/pkg/client"

	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Client Command.",
	Long:  `go-cloud-calculator client 'ip:port'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			fmt.Println("Launching client...")
			var ip_port string = strings.TrimSpace(args[0])
			client.RunClient(ip_port)
		} else {
			fmt.Println("Wrong number of input args")
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
