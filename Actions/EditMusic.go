package actions

import (
	"database/sql"
	"fmt"

	"github.com/manifoldco/promptui"
)

func EditMusic(musicId int) {
	menuOptions := []string{
		"Title",
		"Artist",
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
		case "Title":
			selectTitle(musicId)
		case "Artist":
			selectArtist(musicId)
		case "Exit":
			fmt.Println("Exiting...")
			ListMusic()
		}

		fmt.Println("\nPress Enter to continue...")
	}
}

func selectTitle(musicId int) {
	db, err := sql.Open("sqlite3", "sonata.db")
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	defer db.Close()
	musicNamePrompt := promptui.Prompt{
		Label: "Write the new title of the music",
	}

	musicName, err := musicNamePrompt.Run()

	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	db.Exec("UPDATE music SET title = ? WHERE id = ?", musicName, musicId)
}

func selectArtist(musicId int) {
	db, err := sql.Open("sqlite3", "sonata.db")

	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	defer db.Close()

	musicNamePrompt := promptui.Prompt{
		Label: "Write the new artist of the music",
	}

	musicName, err := musicNamePrompt.Run()

	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	db.Exec("UPDATE music SET artist = ? WHERE id = ?", musicName, musicId)
}
