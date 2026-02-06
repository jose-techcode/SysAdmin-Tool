package cmd

import (
	"fmt"

	"github.com/jose-techcode/SysAdmin-Tool/apiSysinfo"
	"github.com/spf13/cobra"
)

// API REST (GET)

var apiCmd = &cobra.Command{
	Use: "api",
	Short: "Starts the API HTTP server. Localhost.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting HTTP server on port 8000")
		api.Server()
	},
}

// Subcommand registration

func init() {
	rootCmd.AddCommand(apiCmd)
}