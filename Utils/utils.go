package utils

import (
	"fmt"
	"os"
)

func CheckOrCreateFolderLocal() {
	folderName := "local"
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			fmt.Println("err creating folder ", err)
			return
		}
		fmt.Println("Folder local created")
	}
}
