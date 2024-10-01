package models

type Music struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	FilePath string `json:"filepath"`
	Duration int    `json:"duration"`
	ArtistId int    `json:"artistid"`
}
