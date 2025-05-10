package cmd

import (
	"os"
	"fmt"
	"regexp"
)

func changeGrubConfig(changeThis string, changeTo string) {
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

	modifedData := re.ReplaceAllString(dataButWorks, changeThis + "=" + changeTo)

	err = os.WriteFile(grubDefaultFile, []byte(modifedData), 0644)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Println("The operation has been completed successfully.")
}