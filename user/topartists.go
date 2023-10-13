package user

import (
	"encoding/json"
	"fmt"

	"github.com/meinantoyuriawan/spotifyApi/models"
)

func topArtists(AccessToken, term, limit string) models.Artists {
	url := "https://api.spotify.com/v1/me/top/artists"

	s := createTopUserRequests(url, term, limit, AccessToken)

	dataResp := models.Artists{}
	err := json.Unmarshal([]byte(s), &dataResp)
	if err != nil {
		fmt.Println(err)
	}

	return dataResp
}
