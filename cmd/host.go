package cmd

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/spf13/cobra"
)

// Command root

var hostCmd = &cobra.Command{
	Use: "host",
	Short: "Shows kernel and operating system information",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("-----KERNEL-----")

		// Host (Kernel & Operating System)
				
		host_info, err := host.Info()
		if err != nil {
			fmt.Println("Error retrieving information about Kernel and Operating System.", err)
			return
		}

		// Kernel

		fmt.Println("Kernel:", host_info.OS)
	    fmt.Println("Kernel Version:", host_info.KernelVersion)
		fmt.Println("Kernel Architecture:", host_info.KernelArch)
		fmt.Println("")
		fmt.Println("-----OPERATING SYSTEM-----")
        
		// Operating System

		fmt.Println("Hostname:", host_info.Hostname)
		fmt.Println("Plataform:", host_info.Platform)
		fmt.Println("Uptime (minutes):", host_info.Uptime / 60)
		fmt.Println("")
	},
}

// Subcommand registration

func init() {
	rootCmd.AddCommand(hostCmd)
}