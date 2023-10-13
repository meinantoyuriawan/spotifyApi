package models

type Artist struct {
	ArtistName string `json:"name"`
}

type TrackItems struct {
	Artists    []Artist `json:"artists"`
	Href       string   `json:"href"`
	TrackTitle string   `json:"name"`
}

type UserTrack struct {
	Items []TrackItems `json:"items"`
}
