package cmd

import (
	"os"
	"os/exec"
	"fmt"
	"regexp"
	"github.com/spf13/cobra"
	"syscall"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grub-dfct",
	Short: "grub-dfct, a Grub Default File Configuration Tool",
	Long:  "grub-dfct, a Grub Default File Configuration Tool\n" +
		"This program edits /etc/default/grub and can also reconfigure grub.cfg\n",
}
var setTimeout = &cobra.Command{
    Use:   "settimeout",
    Short: "This option sets timeout for grub.",
    Long: "This option sets timeout for grub.\n" +
          "The 30 is delay in seconds, you can set it to anything else.\n" +
          "Remember that you need to apply the changes!",
    
	Run: func(cmd *cobra.Command, args []string) {
		const grubDefaultFile = "/etc/default/grub"

		data, err := os.ReadFile(grubDefaultFile)
		if err != nil { 
			fmt.Println("Error reading:", err)
			return
		}
		dataButWorks := string(data)
		//fmt.Println(string(data))
		fmt.Println(args[0])

		re := regexp.MustCompile(`(?m)^GRUB_TIMEOUT=.*`)

		modifedData := re.ReplaceAllString(dataButWorks, "GRUB_TIMEOUT="+args[0])

		err = os.WriteFile(grubDefaultFile, []byte(modifedData), 0644)
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
		
		fmt.Println("\nTimeout updated.")
    },
}

var updateGrub = &cobra.Command{
	Use:   "update",
    Short: "This option apply's your grub settings.",
    Long: "This option uses grub-mkconfig to reload your settings.\n",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("grub2-mkconfig", "-o", "/boot/grub2/grub.cfg")
		
		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing grub2-mkconfig:", err)
		}
		fmt.Println(string(output))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grub-dfct.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// Check if we have sudo permisions:
	euid := syscall.Geteuid()

	if euid != 0 {
		fmt.Println("This program needs sudo priviliges.")
		os.Exit(1)
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(setTimeout)
	rootCmd.AddCommand(updateGrub)
}


