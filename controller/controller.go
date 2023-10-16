package controller

import (
	"net/http"

	"reflect"

	"github.com/gorilla/mux"
	"github.com/meinantoyuriawan/spotifyApi/helper"
	"github.com/meinantoyuriawan/spotifyApi/models"
	"github.com/meinantoyuriawan/spotifyApi/recommendations"
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

func GetUserTop(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	reqType := params["type"]
	term := params["term"]
	limit := params["limit"]

	createUserTopLogic(reqType, term, limit, w)
}

func GetUserTopDefault(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	reqType := params["type"]
	term := "medium"
	limit := "10"

	createUserTopLogic(reqType, term, limit, w)
}

func createUserTopLogic(reqType, term, limit string, w http.ResponseWriter) {
	if reqType == "tracks" {
		TopTracks := user.GetUserTopTracks(term, limit)
		if reflect.DeepEqual(TopTracks, models.UserTrack{}) {
			helper.CreateErrorResponse(w, "User not logged in", http.StatusBadRequest)
		} else {
			helper.ResponseJSON(w, http.StatusOK, TopTracks)
		}
	} else if reqType == "artists" {
		TopArtists := user.GetUserTopArtists(term, limit)
		if reflect.DeepEqual(TopArtists, models.Artists{}) {
			helper.CreateErrorResponse(w, "User not logged in", http.StatusBadRequest)
		} else {
			helper.ResponseJSON(w, http.StatusOK, TopArtists)
		}
	} else {
		helper.CreateErrorResponse(w, "Wrong params", http.StatusBadRequest)
	}
}

func GetRecommendationArtist(w http.ResponseWriter, r *http.Request) {
	UserRec := recommendations.GetTrackRecommendations()
	if reflect.DeepEqual(UserRec, models.BySeedsRecommendation{}) {
		helper.CreateErrorResponse(w, "User not logged in", http.StatusBadRequest)
	} else {
		helper.ResponseJSON(w, http.StatusOK, UserRec)
	}
}

func DisplayError(w http.ResponseWriter, r *http.Request) {
	helper.CreateErrorResponse(w, "state_mismatch", http.StatusBadRequest)
}
