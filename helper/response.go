package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {

	response, _ := json.Marshal(payload)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CreateErrorResponse(w http.ResponseWriter, message string, code int) {
	response := map[string]string{"message": message}
	ResponseJSON(w, code, response)
}
