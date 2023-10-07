package controller

import (
	"net/http"

	"github.com/meinantoyuriawan/spotifyApi/helper"
	spotifyauth "github.com/meinantoyuriawan/spotifyApi/spotifyAuth"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	token, err := spotifyauth.GetToken()

	if err != nil {
		helper.CreateErrorResponse(w, "token unavailable", http.StatusBadRequest)
	}

	helper.ResponseJSON(w, http.StatusOK, token)

}
