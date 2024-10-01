package actions

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func AddMusic() {
	musicNamePrompt := promptui.Prompt{
		Label: "Write the music name",
	}

	musicName, err := musicNamePrompt.Run()

	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	musicArtistPrompt := promptui.Prompt{
		Label: "Write the music name",
	}

	musicArtist, err := musicArtistPrompt.Run()

	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Println(musicArtist, musicName)

}

func EditMusic() {
	fmt.Println("you choose add music")

}

func ListMusic() {
	fmt.Println("you choose add music")

}

func DeleteMusic() {
	fmt.Println("you choose add music")

}

func GetMusicByName() {
	fmt.Println("you choose add music")

}

func GetMusicByArtist() {
	fmt.Println("you choose add music")

}

func GetMusicByPlaylist() {
	fmt.Println("you choose add music")

}
