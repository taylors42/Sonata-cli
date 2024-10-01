package main

import (
	"fmt"
	"os"

	actions "sonata-cli/Actions"
	"sonata-cli/database"

	"github.com/manifoldco/promptui"
)

func main() {
	database.ConnectToDb()

	// Define the menu options
	menuOptions := []string{
		"Music List",
		"Playlists",
		"Add Music",
		"Edit Music",
		"Delete Music",
		"Exit",
	}

	// Create a prompt for the menu
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
			actions.ListMusic()
		case "Playlists":
			actions.ListPlaylist()
		case "Add Music":
			actions.AddMusic()
		case "Edit Music":
			actions.EditMusic()
		case "Delete Music":
			actions.DeleteMusic()
		case "Exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()
	}
}
