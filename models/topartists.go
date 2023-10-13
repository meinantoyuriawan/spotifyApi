package models

type ArtistUrl struct {
	Url string `json:"spotify"`
}

type ArtistItems struct {
	Genres    []string  `json:"genres"`
	Artists   string    `json:"name"`
	ArtistUrl ArtistUrl `json:"external_urls"`
}

type Artists struct {
	Items []ArtistItems `json:"items"`
}
