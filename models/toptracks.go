package models

type Artist struct {
	ArtistName string `json:"name"`
}

type TrackLink struct {
	Href string `json:"spotify"`
}

type TrackItems struct {
	Artists    []Artist  `json:"artists"`
	Href       string    `json:"href"`
	Id         string    `json:"id"`
	TrackLink  TrackLink `json:"external_urls"`
	TrackTitle string    `json:"name"`
}

type UserTrack struct {
	Items []TrackItems `json:"items"`
}
