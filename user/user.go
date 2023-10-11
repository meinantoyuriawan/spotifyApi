package user

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/meinantoyuriawan/spotifyApi/helper"
)

func GetUserProfile() {

	isLogin, AccessToken := isLogin()

	if !isLogin {
		fmt.Println("not logged in")
	}

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

	fmt.Println(s)
}

// todo:
// error handling user top tracks
// custom time range, limit and offset
func GetUserTopTracks(w http.ResponseWriter) {

	isLogin, AccessToken := isLogin()

	if !isLogin {
		fmt.Println("not logged in")
	}

	userTracks := topTracks(AccessToken)

	helper.ResponseJSON(w, http.StatusOK, userTracks)
}

//todo:
// Get User Top Artist

func isLogin() (bool, string) {
	//get token
	AccToken := helper.ReadToken()

	if AccToken == "" {
		fmt.Println("user not logged in, go to /login")
		return false, ""
	}

	return true, AccToken

}
