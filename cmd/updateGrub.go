package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "This option applies your GRUB settings.",
	Long:  "This option uses grub-mkconfig to reload your settings.\n",
	Run: func(cmd *cobra.Command, args []string) {

		//Check for sudo permission:
		euid := syscall.Geteuid()

		if euid != 0 {
			fmt.Println("This program needs sudo priviliges.")
			os.Exit(1)
		}

		command := exec.Command("grub2-mkconfig", "-o", "/boot/grub2/grub.cfg")

		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing grub2-mkconfig:", err)
		}
		fmt.Println("Output of grub mkconfig:")
		fmt.Println(string(output))
	},
}
