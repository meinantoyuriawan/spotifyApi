package models

type RecommendedLink struct {
	Href string `json:"spotify"`
}

type RecommendedTrack struct {
	RecommendedArtists []Artist        `json:"artists"`
	TrackLink          RecommendedLink `json:"external_urls"`
	TrackTitle         string          `json:"name"`
}

type UserRecommendation struct {
	Tracks []RecommendedTrack `json:"tracks"`
}

type BySeedsRecommendation struct {
	ByTracks  UserRecommendation `json:"by_tracks"`
	ByArtists UserRecommendation `json:"by_artists"`
	ByGenres  UserRecommendation `json:"by_genres"`
	ByRandom  UserRecommendation `json:"by_random"`
}
