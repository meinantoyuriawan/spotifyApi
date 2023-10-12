package user

import (
	"fmt"
	"net/http"

	"github.com/meinantoyuriawan/spotifyApi/helper"
)

// todo:
// error handling user
func GetUserProfile(w http.ResponseWriter) {

	isLogin, AccessToken := isLogin()

	if !isLogin {
		fmt.Println(AccessToken)
	} else {
		Profile := userProfile(AccessToken)

		helper.ResponseJSON(w, http.StatusOK, Profile)
	}
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
		// fmt.Println("user not logged in, go to /login")
		return false, "user not logged in, go to /login"
	}

	return true, AccToken

}
