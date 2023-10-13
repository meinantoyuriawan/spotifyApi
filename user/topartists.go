package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/meinantoyuriawan/spotifyApi/models"
)

func topArtists(AccessToken, term, limit string) models.Artists {
	url := "https://api.spotify.com/v1/me/top/artists"

	Authorization := "Bearer " + AccessToken

	req, err := http.NewRequest("GET", url, nil)

	if term == "short" {
		term = "short_term"
	} else {
		term = "medium_term"
	}

	q := req.URL.Query()
	q.Add("time_range", term)
	q.Add("limit", limit)
	q.Add("offset", "0")

	req.URL.RawQuery = q.Encode()

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", Authorization)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		// return "", err
		fmt.Println("err resp")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// return "", err
		fmt.Println("err Read")
	}

	s := string(body)

	dataResp := models.Artists{}
	err = json.Unmarshal([]byte(s), &dataResp)
	if err != nil {
		fmt.Println(err)
	}

	return dataResp
}
