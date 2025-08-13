package cmd

import(
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// Command root

var usbCmd = &cobra.Command{
	Use: "usb",
	Short: "Exibe informações de dispositivos USB.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("-----USB-----")

		// USB

		usb_info, err := exec.Command("lsusb").Output()
		if err != nil {
			fmt.Println("Error retrieving information about usb devices.", err)
			return
		}

		usb_text := strings.TrimSpace(string(usb_info))

		fmt.Println(usb_text)
		fmt.Println("")
	},
}

// Subcommand registration

func init() {
	rootCmd.AddCommand(usbCmd)
}