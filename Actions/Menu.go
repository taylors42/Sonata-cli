package actions

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func ShowMenuList() {
	menuOptions := []string{
		"Music List",
		"Playlists",
		"Add Music",
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
		case "Music List":
			ListMusic()
		case "Playlists":
			ListPlaylist()
		case "Add Music":
			AddMusic()
		case "Exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		fmt.Println("\nPress Enter to continue...")
	}
}
