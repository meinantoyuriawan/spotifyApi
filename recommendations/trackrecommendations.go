package recommendations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/meinantoyuriawan/spotifyApi/helper"
	"github.com/meinantoyuriawan/spotifyApi/models"
)

func GetTrackRecommendations() models.BySeedsRecommendation {

	url := "https://api.spotify.com/v1/recommendations"

	//get token
	AccessToken := helper.ReadToken()

	isLogin := isLogin(AccessToken)

	recommendations := models.BySeedsRecommendation{}

	if !isLogin {
		// returning empty Recommendation
		fmt.Println("not loggedin")
		return recommendations
	}

	// else get the seeds for recommendation

	trackSeeds, artistsSeeds, genresSeeds, randomSeeds, a := GetSeeds()

	market := getUserCountry()

	byTrack := createRecRequest(url, AccessToken, market, "track", trackSeeds)
	byArtist := createRecRequest(url, AccessToken, market, "artist", artistsSeeds)
	byGenres := createRecRequest(url, AccessToken, market, "genre", genresSeeds)
	byRandom := createRandomRequest(url, AccessToken, market, randomSeeds, a)

	recommendations.ByTracks = byTrack
	recommendations.ByArtists = byArtist
	recommendations.ByGenres = byGenres
	recommendations.ByRandom = byRandom

	return recommendations
}

func isLogin(AccToken string) bool {
	return AccToken != ""
}

func clientHit(req *http.Request) models.UserRecommendation {

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

	dataResp := models.UserRecommendation{}
	err = json.Unmarshal([]byte(s), &dataResp)
	if err != nil {
		fmt.Println(err)
	}

	return dataResp
}

func createRandomRequest(url, AccessToken, market string, seeds []string, a [3]int) models.UserRecommendation {

	Authorization := "Bearer " + AccessToken

	req, err := http.NewRequest("GET", url, nil)

	strParams := ""

	for _, vals := range seeds[:4] {
		strParams += vals + ","
	}

	strParams += seeds[4]

	q := req.URL.Query()
	q.Add("limit", "10")
	q.Add("market", market)

	trackAmount := a[0]
	artistAmount := a[1]
	genresAmount := a[2]

	trackParams := ""
	artistParams := ""
	genreParams := ""

	for i := 0; i < trackAmount; i++ {
		trackParams += seeds[i]
	}

	for i := trackAmount; i < artistAmount; i++ {
		artistParams += seeds[i]
	}

	for i := artistAmount; i < genresAmount; i++ {
		genreParams += seeds[i]
	}

	q.Add("seed_tracks", trackParams)
	q.Add("seed_artists", artistParams)
	q.Add("seed_genres", genreParams)

	req.URL.RawQuery = q.Encode()

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", Authorization)

	dataResp := clientHit(req)

	return dataResp
}

func createRecRequest(url, AccessToken, market, identifier string, seeds []string) models.UserRecommendation {

	Authorization := "Bearer " + AccessToken

	req, err := http.NewRequest("GET", url, nil)

	strParams := ""

	for _, vals := range seeds[:4] {
		strParams += vals + ","
	}

	strParams += seeds[4]

	q := req.URL.Query()
	q.Add("limit", "10")
	q.Add("market", market)

	if identifier == "track" {
		q.Add("seed_tracks", strParams)
	} else if identifier == "artist" {
		q.Add("seed_artists", strParams)
	} else if identifier == "genre" {
		q.Add("seed_genres", strParams)
	}

	req.URL.RawQuery = q.Encode()

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", Authorization)

	dataResp := clientHit(req)

	return dataResp
}
