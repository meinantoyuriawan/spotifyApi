package spotifyauth

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TokenResponse struct {
	Token string `json:"access_token"`
	Type  string `json:"token_type"`
	Time  int    `json:"expires_in"`
}

func GetToken() (string, error) {

	url := "https://accounts.spotify.com/api/token"

	CLIENT_ID := ""
	CLIENT_SECRET := ""

	strSecret := CLIENT_ID + ":" + CLIENT_SECRET

	sEnc := b64.StdEncoding.EncodeToString([]byte(strSecret))

	Authorization := "Basic " + string(sEnc)

	payload := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println("hay")
		return "", err
	}

	req.Header.Add("Authorization", Authorization)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("hay")
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("hay")
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
