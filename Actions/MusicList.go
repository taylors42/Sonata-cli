package actions

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

type Music struct {
	ID     int
	Name   string
	Artist string
}

func ListMusic() {
	db, err := sql.Open("sqlite3", "sonata.db")

	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	defer db.Close()

	listMusicQuery := `
		SELECT [id], [title], [artist] FROM music;
	`
	rows, err := db.Query(listMusicQuery)

	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	defer rows.Close()

	var musics []Music

	for rows.Next() {
		var music Music
		if err := rows.Scan(&music.ID, &music.Name, &music.Artist); err != nil {
			log.Fatal(err)
		}
		musics = append(musics, music)
	}

	if len(musics) == 0 {
		fmt.Println("No music found")
		ShowMenuList()
		return
	}

	var items []string
	var musicId int
	for _, music := range musics {
		items = append(items, fmt.Sprintf("%d. %s - %s", music.ID, music.Name, music.Artist))
		musicId = music.ID
	}

	prompt := promptui.Select{
		Label: "Select Music",
		Items: items,
	}
	_, _, err = prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	SelectAction(musicId)
}
