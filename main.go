package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meinantoyuriawan/spotifyApi/controller"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/test-cred", controller.GetToken).Methods("GET")

	r.HandleFunc("/login", controller.Login).Methods("GET")

	r.HandleFunc("/login-error", controller.DisplayError).Methods("GET")

	r.HandleFunc("/callback", controller.CallbackLogin).Methods("GET")

	r.HandleFunc("/getuser", controller.GetUser).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	// http.Handle("/", r)
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))
}
