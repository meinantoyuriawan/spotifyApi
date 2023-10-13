package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/meinantoyuriawan/spotifyApi/models"
)

func userProfile(AccessToken string) models.Profile {
	url := "https://api.spotify.com/v1/me"

	Authorization := "Bearer " + AccessToken

	req, err := http.NewRequest("GET", url, nil)

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

	dataResp := models.Profile{}
	err = json.Unmarshal([]byte(s), &dataResp)
	if err != nil {
		fmt.Println(err)
	}

	return dataResp
}
