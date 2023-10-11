package spotifyauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/meinantoyuriawan/spotifyApi/helper"
)

type TokenResponse struct {
	Token string `json:"access_token"`
	Type  string `json:"token_type"`
	Time  int    `json:"expires_in"`
}

type CallBackLoginResponse struct {
	AccToken     string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	Time         int    `json:"expires_in"`
	TokenRefresh string `json:"refresh_token"`
}

func GetTokenClientCred() (string, error) {

	url := "https://accounts.spotify.com/api/token"

	Authorization := helper.GenerateBasicToken()

	payload := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", Authorization)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	s := string(body)

	data := TokenResponse{}
	err = json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	token := data.Token

	return token, nil
}

func TriggerAuthByCode() (string, error) {
	req, err := http.NewRequest("GET", "https://accounts.spotify.com/authorize", nil)

	if err != nil {
		return "", err
	}

	url := redirectAuthByCode(req)

	return url, nil
}

func redirectAuthByCode(r *http.Request) string {
	redirect_uri := "http://localhost:8080/callback"
	state := "abc123"
	scope := "user-read-private user-read-email user-top-read user-read-recently-played"
	id := " " //client id

	q := r.URL.Query()
	q.Add("response_type", "code")
	q.Add("client_id", id)
	q.Add("scope", scope)
	q.Add("redirect_uri", redirect_uri)
	q.Add("state", state)

	r.URL.RawQuery = q.Encode()

	return r.URL.String()

	// http.Redirect(w, r, r.URL.String(), http.StatusSeeOther)
}

func CallbackAuthByCode(r *http.Request) (string, error) {

	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if state == "" {
		// helper.CreateErrorResponse(w, "state_mismatch", http.StatusBadRequest)
		return "", errors.New("state_mismatch")
	}

	AccToken, err := getTokenAuthByCode(code)
	// fmt.Println(AccToken)

	if err != nil {
		return "", err
	}

	return AccToken, nil

}

func getTokenAuthByCode(code string) (string, error) {
	redirect_uri := "http://localhost:8080/callback"

	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", redirect_uri)
	data.Set("grant_type", "authorization_code")

	encodedData := data.Encode()

	url := "https://accounts.spotify.com/api/token"

	req, err := http.NewRequest("POST", url, strings.NewReader(encodedData))

	if err != nil {
		return "", err
	}

	Authorization := helper.GenerateBasicToken()

	req.Header.Add("Authorization", Authorization)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	s := string(body)

	dataResp := CallBackLoginResponse{}
	err = json.Unmarshal([]byte(s), &dataResp)
	if err != nil {
		return "", err
	}

	// fmt.Println(string(dataResp.AccToken))
	return dataResp.AccToken, nil
}
