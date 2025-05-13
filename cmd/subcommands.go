package cmd

import (
	"os/exec"
	"fmt"
	"github.com/spf13/cobra"
)

var setTimeout = &cobra.Command{
    Use:   "settimeout",
    Short: "This option sets timeout for grub.",
    Long: "This option sets timeout for grub.\n" +
          "The 30 is delay in seconds, you can set it to anything else.\n" +
          "Remember that you need to apply the changes!",
    
	Run: func(cmd *cobra.Command, args []string) {
		changeGrubConfig("GRUB_TIMEOUT", args[0])
    },
}

var updateGrub = &cobra.Command{
	Use:   "update",
    Short: "This option apply's your grub settings.",
    Long: "This option uses grub-mkconfig to reload your settings.\n",
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
		fmt.Println(string(output))
	},
}