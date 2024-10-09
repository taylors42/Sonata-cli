package actions

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

func DeleteMusic(musicId int) {
	menuOptions := []string{
		"Yes",
		"No",
	}

	prompt := promptui.Select{
		Label: "Are you sure you want to delete this music?",
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
		case "Yes":
			RemoveMusic(musicId)
		case "No":
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}
}

func RemoveMusic(musicId int) {
	db, err := sql.Open("sqlite3", "sonata.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	removeMusicQuery := "DELETE FROM music WHERE id = ?"

	_, err = db.Exec(removeMusicQuery, musicId)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Music removed successfully")
	ListMusic()
}
