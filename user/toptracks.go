package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Artist struct {
	ArtistName string `json:"name"`
}

type TrackItems struct {
	Artists    []Artist `json:"artists"`
	Href       string   `json:"href"`
	TrackTitle string   `json:"name"`
}

type UserTrack struct {
	Items []TrackItems `json:"items"`
}

func topTracks(AccessToken string) UserTrack {
	url := "https://api.spotify.com/v1/me/top/tracks"

	Authorization := "Bearer " + AccessToken

	req, err := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("time_range", "short_term")
	q.Add("limit", "10")
	q.Add("offset", "0")

	req.URL.RawQuery = q.Encode()

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", Authorization)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		// return "", err
		fmt.Println("err resp")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// return "", err
		fmt.Println("err Read")
	}

	s := string(body)

	dataResp := UserTrack{}
	err = json.Unmarshal([]byte(s), &dataResp)
	if err != nil {
		fmt.Println(err)
	}

	TopUserTracksResponse := UserTrack{}
	TopTrackResponse := TrackItems{}

	for _, dataItems := range dataResp.Items {
		TopTrackResponse.Artists = dataItems.Artists
		TopTrackResponse.Href = dataItems.Href
		TopTrackResponse.TrackTitle = dataItems.TrackTitle

		TopUserTracksResponse.Items = append(TopUserTracksResponse.Items, TopTrackResponse)
	}

	response, _ := json.Marshal(TopUserTracksResponse)

	fmt.Println(string(response))

	return TopUserTracksResponse
}
