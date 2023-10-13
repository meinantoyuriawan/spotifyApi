package user

import (
	"encoding/json"
	"fmt"

	"github.com/meinantoyuriawan/spotifyApi/models"
)

func topTracks(AccessToken, term, limit string) models.UserTrack {
	url := "https://api.spotify.com/v1/me/top/tracks"

	s := createTopUserRequests(url, term, limit, AccessToken)
	dataResp := models.UserTrack{}
	err := json.Unmarshal([]byte(s), &dataResp)
	if err != nil {
		fmt.Println(err)
	}

	return dataResp
}
