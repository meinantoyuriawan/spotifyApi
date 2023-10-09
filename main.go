package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meinantoyuriawan/spotifyApi/controller"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/test-cred", controller.GetToken).Methods("GET")

	r.HandleFunc("/login", controller.Login).Methods("GET")

	r.HandleFunc("/callback", controller.CallbackLogin).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
