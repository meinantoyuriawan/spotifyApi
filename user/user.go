package user

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/meinantoyuriawan/spotifyApi/helper"
)

func GetUserProfile() {

	//get token
	AccToken := helper.ReadToken()

	if AccToken == "" {
		fmt.Println("user not logged in, go to /login")
	}

	url := "https://api.spotify.com/v1/me"

	Authorization := "Bearer " + AccToken

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("user not logged in")
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

	fmt.Println(s)
}
