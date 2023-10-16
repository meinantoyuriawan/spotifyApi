package recommendations

import (
	"encoding/json"
	"fmt"

	"github.com/meinantoyuriawan/spotifyApi/models"
)

func getUserCountry() string {
	url := "http://localhost:8080/get-user"

	s := hitSeedsApi(url)

	dataResp := models.Profile{}
	err := json.Unmarshal([]byte(s), &dataResp)
	if err != nil {
		fmt.Println(err)
	}

	return dataResp.Country
}
