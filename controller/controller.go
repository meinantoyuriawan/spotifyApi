package controller

import (
	"fmt"
	"net/http"

	"github.com/meinantoyuriawan/spotifyApi/helper"
	spotifyauth "github.com/meinantoyuriawan/spotifyApi/spotifyAuth"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	token, err := spotifyauth.GetTokenClientCred()

	if err != nil {
		helper.CreateErrorResponse(w, "token unavailable", http.StatusBadRequest)
	}

	helper.ResponseJSON(w, http.StatusOK, token)

}

func Login(w http.ResponseWriter, r *http.Request) {
	url, err := spotifyauth.TriggerAuthByCode()

	if err != nil {
		helper.CreateErrorResponse(w, err.Error(), http.StatusBadRequest)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func CallbackLogin(w http.ResponseWriter, r *http.Request) {
	AccToken, err := spotifyauth.CallbackAuthByCode(r)

	if err != nil {
		helper.CreateErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	helper.ResponseJSON(w, http.StatusOK, AccToken)

	fmt.Println(AccToken)
}
