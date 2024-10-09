package actions

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func SelectAction(musicId int) {
	menuOptions := []string{
		"Play",
		"Edit",
		"Delete",
		"Exit",
	}

	prompt := promptui.Select{
		Label: "Select Option",
		Items: menuOptions,
		Size:  len(menuOptions),
	}

	for {
		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "Play":
			PlayMusic()
		case "Edit":
			EditMusic(musicId)
		case "Delete":
			DeleteMusic(musicId)
		case "Exit":
			fmt.Println("Exiting...")
			ShowMenuList()
		}

		fmt.Println("\nPress Enter to continue...")
	}
}
