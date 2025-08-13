package cmd

import (
	"fmt"

	"github.com/jose-techcode/CLI_Luasys/api"
	"github.com/spf13/cobra"
)

// API REST (GET)

var apiCmd = &cobra.Command{
	Use: "api",
	Short: "Inicia servidor HTTP da API. Localhost.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Iniciando servidor HTTP na porta 8000...")
		api.Server()
	},
}

// Subcommand registration

func init() {
	rootCmd.AddCommand(apiCmd)
}