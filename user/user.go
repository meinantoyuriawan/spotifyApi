package user

import (
	"github.com/meinantoyuriawan/spotifyApi/helper"
	"github.com/meinantoyuriawan/spotifyApi/models"
)

func GetUserProfile() models.Profile {

	//get token
	AccessToken := helper.ReadToken()

	isLogin := isLogin(AccessToken)

	Profile := models.Profile{}

	if !isLogin {
		// returning empty Profile
		return Profile
	}

	Profile = userProfile(AccessToken)

	return Profile
}

// todo:
// custom time range, limit and offset
func GetUserTopTracks() models.UserTrack {

	//get token
	AccessToken := helper.ReadToken()

	isLogin := isLogin(AccessToken)

	UserTracks := models.UserTrack{}

	if !isLogin {
		// returning empty UserTracks
		return UserTracks
	}

	UserTracks = topTracks(AccessToken)

	return UserTracks
}

//todo:
// Get User Top Artist

func isLogin(AccToken string) bool {
	return AccToken != ""
}
