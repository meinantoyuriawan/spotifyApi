package controller

import (
	"net/http"

	"reflect"

	"github.com/gorilla/mux"
	"github.com/meinantoyuriawan/spotifyApi/helper"
	"github.com/meinantoyuriawan/spotifyApi/models"
	spotifyauth "github.com/meinantoyuriawan/spotifyApi/spotifyAuth"
	"github.com/meinantoyuriawan/spotifyApi/user"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	token, err := spotifyauth.GetTokenClientCred()

	if err != nil || token == "" {
		helper.CreateErrorResponse(w, "token unavailable", http.StatusBadRequest)
	}

	helper.ResponseJSON(w, http.StatusOK, token)

}

func Login(w http.ResponseWriter, r *http.Request) {
	url, err := spotifyauth.TriggerAuthByCode()

	if err != nil {
		http.Redirect(w, r, "/login-error", http.StatusSeeOther)
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func CallbackLogin(w http.ResponseWriter, r *http.Request) {
	AccToken, err := spotifyauth.CallbackAuthByCode(r)

	if err != nil || AccToken == "" {
		helper.CreateErrorResponse(w, "invalid token", http.StatusBadRequest)
	} else {
		helper.WriteToken(AccToken)
		helper.ResponseJSON(w, http.StatusOK, AccToken)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	UserProfile := user.GetUserProfile()
	if reflect.DeepEqual(UserProfile, models.Profile{}) {
		// fmt.Println("halo")
		helper.CreateErrorResponse(w, "User not logged in", http.StatusBadRequest)
	} else {
		helper.ResponseJSON(w, http.StatusOK, UserProfile)
	}
}

func GetTopArtist(w http.ResponseWriter, r *http.Request) {

	// user.GetUserTop(w)
}

func GetTopTracksDefault(w http.ResponseWriter, r *http.Request) {

	term := "medium"
	limit := "10"

	TopTracks := user.GetUserTopTracks(term, limit)
	if reflect.DeepEqual(TopTracks, models.UserTrack{}) {
		helper.CreateErrorResponse(w, "User not logged in", http.StatusBadRequest)
	} else {
		helper.ResponseJSON(w, http.StatusOK, TopTracks)
	}
}

func GetTopTracks(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	term := params["term"]
	limit := params["limit"]

	TopTracks := user.GetUserTopTracks(term, limit)
	if reflect.DeepEqual(TopTracks, models.UserTrack{}) {
		helper.CreateErrorResponse(w, "User not logged in", http.StatusBadRequest)
	} else {
		helper.ResponseJSON(w, http.StatusOK, TopTracks)
	}
}

func DisplayError(w http.ResponseWriter, r *http.Request) {
	helper.CreateErrorResponse(w, "state_mismatch", http.StatusBadRequest)
}
