package user

import (
	"fmt"
	"net/http"

	"github.com/meinantoyuriawan/spotifyApi/helper"
)

// todo:
// error handling user
func GetUserProfile(w http.ResponseWriter) {

	//get token
	AccessToken := helper.ReadToken()

	isLogin := isLogin(AccessToken)

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

	//get token
	AccessToken := helper.ReadToken()

	isLogin := isLogin(AccessToken)

	if !isLogin {
		fmt.Println("not logged in")
	}

	userTracks := topTracks(AccessToken)

	helper.ResponseJSON(w, http.StatusOK, userTracks)
}

//todo:
// Get User Top Artist

func isLogin(AccToken string) bool {

	// if AccToken == "" {
	// 	return false
	// }

	// return true

	return AccToken != ""

}
