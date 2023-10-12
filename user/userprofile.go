package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Url struct {
	UrlProfile string `json:"spotify"`
}

type Image struct {
	ImageUrl string `json:"url"`
}

type Profile struct {
	Name   string  `json:"display_name"`
	Urls   Url     `json:"external_urls"`
	Images []Image `json:"images"`
	Email  string  `json:"email"`
}

func userProfile(AccessToken string) Profile {
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

	dataResp := Profile{}
	err = json.Unmarshal([]byte(s), &dataResp)
	if err != nil {
		fmt.Println(err)
	}

	return dataResp
}
