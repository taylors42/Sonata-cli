package models

type Playlist struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
