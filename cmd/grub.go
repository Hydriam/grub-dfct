package cmd

import (
	"fmt"
	"os"
	"regexp"
	"syscall"
)

func changeGrubConfig(changeThis string, changeTo string) {
	//Check for sudo permission:
	euid := syscall.Geteuid()

	if euid != 0 {
		fmt.Println("This program needs sudo privileges.")
		os.Exit(1)
	}

	const grubDefaultFile = "/etc/default/grub"
	data, err := os.ReadFile(grubDefaultFile)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	dataButWorks := string(data)

	idkitworks := fmt.Sprintf(`(?m)^%s=.*`, changeThis)
	re := regexp.MustCompile(idkitworks)

	//re := regexp.MustCompile(`(?m)^%s=.*`, changeThis)

	modifedData := re.ReplaceAllString(dataButWorks, changeThis+"="+changeTo)

	err = os.WriteFile(grubDefaultFile, []byte(modifedData), 0644)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Println("The operation has been completed successfully.")
}
