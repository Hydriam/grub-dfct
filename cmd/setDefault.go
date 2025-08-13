package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var setDefaultCmd = &cobra.Command{
	Use:   "set-default-entry",
	Short: "This option sets default entry for grub.",
	Long: `This option sets default entry for grub.
You can set it to a number or to "saved".
Remember that you need to apply the changes!`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: You should specify the default entry.")
			os.Exit(1)
		}

		changeTo := args[0]

		if changeTo != "saved" {
			num, err := strconv.Atoi(changeTo)
			if err != nil {
				fmt.Println("Error: Argument Not Valid.")
				os.Exit(1)
			}
			// Grub uses index from 0, if we want to make entry number 1 the first entry from grub we need to decrease by one
			num--
			changeTo = strconv.Itoa(num)
		}

		fmt.Println("Changing grub default entry to " + changeTo)
		changeGrubConfig("GRUB_DEFAULT", changeTo)
	},
}
