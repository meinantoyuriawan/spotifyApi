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

func GetUserTopTracks(term, limit string) models.UserTrack {

	//get token
	AccessToken := helper.ReadToken()

	isLogin := isLogin(AccessToken)

	UserTracks := models.UserTrack{}

	if !isLogin {
		// returning empty UserTracks
		return UserTracks
	}
	UserTracks = topTracks(AccessToken, term, limit)

	return UserTracks
}

func GetUserTopArtists(term, limit string) models.Artists {

	//get token
	AccessToken := helper.ReadToken()

	isLogin := isLogin(AccessToken)

	Artists := models.Artists{}

	if !isLogin {
		// returning empty UserTracks
		return Artists
	}
	Artists = topArtists(AccessToken, term, limit)

	return Artists
}

func isLogin(AccToken string) bool {
	return AccToken != ""
}
