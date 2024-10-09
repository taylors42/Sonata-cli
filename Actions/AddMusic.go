package actions

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func AddMusic() {
	db, err := sql.Open("sqlite3", "sonata.db")

	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	defer db.Close()

	musicNamePrompt := promptui.Prompt{
		Label: "Write the music name",
	}

	musicName, err := musicNamePrompt.Run()

	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	musicArtistPrompt := promptui.Prompt{
		Label: "Write name of the artist",
	}

	musicArtist, err := musicArtistPrompt.Run()
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}
	folderName := "./local"

	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		fmt.Println("Folder 'local' does not exist.")
		return
	}

	var mp3Files []string

	err = filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {
			mp3Files = append(mp3Files, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error while reading the folder:", err)
		return
	}

	if len(mp3Files) == 0 {
		fmt.Println("No .mp3 files found in 'local' folder.")
		return
	}

	getMusicPrompt := promptui.Select{
		Label: "Select an MP3 file",
		Items: mp3Files,
	}

	_, selectedFile, err := getMusicPrompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You selected: %s\n", selectedFile)

	addMusicQuery := `
		INSERT INTO music (title, filepath, artist) 
		VALUES (?,?,?);
	`
	_, err = db.Exec(addMusicQuery, musicName, selectedFile, musicArtist)
	if err != nil {
		fmt.Println("Error while adding music:", err)
		return
	}
	fmt.Println("Music added successfully")
}
