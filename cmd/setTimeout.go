package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var setTimeoutCmd = &cobra.Command{
	Use:   "set-timeout",
	Short: "This option sets timeout for grub.",
	Long: `This option sets timeout for grub." +
"The 30 is delay in seconds, you can set it to anything else." +
"Remember that you need to apply the changes!`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: You should specify the timeout value.")
			os.Exit(1)
		}
		fmt.Println("Changing grub timeout to " + args[0])
		changeGrubConfig("GRUB_TIMEOUT", args[0])
	},
}
