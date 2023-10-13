package controller

import (
	"net/http"

	"github.com/meinantoyuriawan/spotifyApi/helper"
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

	user.GetUserProfile(w)
}

func GetTopArtist(w http.ResponseWriter, r *http.Request) {

	// user.GetUserTop(w)
}

func GetTopTracks(w http.ResponseWriter, r *http.Request) {

	user.GetUserTopTracks(w)
}

func DisplayError(w http.ResponseWriter, r *http.Request) {
	helper.CreateErrorResponse(w, "state_mismatch", http.StatusBadRequest)
}
