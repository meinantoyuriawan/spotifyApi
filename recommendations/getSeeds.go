package recommendations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"

	"github.com/meinantoyuriawan/spotifyApi/models"
)

func GetSeeds() ([]string, []string, []string, []string, [3]int) {
	trackSeeds := getTrackSeeds()

	artistsSeeds, genresSeeds := getArtistsandGenresSeeds()

	randomSeeds, a := surpriseMe(trackSeeds, artistsSeeds, genresSeeds)

	return trackSeeds, artistsSeeds, genresSeeds, randomSeeds, a
}

func surpriseMe(trackSeeds, artistsSeeds, genresSeeds []string) ([]string, [3]int) {
	var a [3]int

	for i := 1; i < 3; i++ {
		a[i] = 1
	}

	for i := 1; i < 4; i++ {
		index := rand.Intn(3)
		a[index] += 1
	}

	var slicesRandom []string

	trackAmount := a[0]
	artistAmount := a[1]
	genresAmount := a[2]

	// track seeds
	for i := 0; i < trackAmount; i++ {
		randomIndex := rand.Intn(len(trackSeeds))
		slicesRandom = append(slicesRandom, trackSeeds[randomIndex])
	}

	for i := 0; i < artistAmount; i++ {
		randomIndex := rand.Intn(len(artistsSeeds))
		slicesRandom = append(slicesRandom, artistsSeeds[randomIndex])
	}

	for i := 0; i < genresAmount; i++ {
		randomIndex := rand.Intn(len(genresSeeds))
		slicesRandom = append(slicesRandom, genresSeeds[randomIndex])
	}
	return slicesRandom, a
}

func getTrackSeeds() []string {
	// hit /get-top/tracks/short/5
	url := "http://localhost:8080/get-top/tracks/short/5"

	s := hitSeedsApi(url)

	dataResp := models.UserTrack{}
	err := json.Unmarshal([]byte(s), &dataResp)
	if err != nil {
		fmt.Println(err)
	}

	// get top 5 tracks seeds
	var trackSeeds []string

	for _, seeds := range dataResp.Items {
		trackSeeds = append(trackSeeds, seeds.Id)
	}

	return trackSeeds
}

func getArtistsandGenresSeeds() ([]string, []string) {
	// hit /get-top/artists/short/5
	url := "http://localhost:8080/get-top/artists/short/10"

	s := hitSeedsApi(url)

	// get top 5 artists
	dataResp := models.Artists{}
	err := json.Unmarshal([]byte(s), &dataResp)
	if err != nil {
		fmt.Println(err)
	}

	var artistSeeds []string

	for _, seeds := range dataResp.Items[:5] {
		artistSeeds = append(artistSeeds, seeds.Id)
	}

	// get top 5 genres

	hashmap := make(map[string]int)

	for _, seeds := range dataResp.Items {
		for _, genres := range seeds.Genres {

			_, ok := hashmap[genres]
			if ok {
				hashmap[genres] += 1
			} else {
				hashmap[genres] = 1
			}
		}
	}

	vec := mapToSlice(hashmap)
	sort.Slice(vec, func(i, j int) bool {
		// value is different - sort by vals
		if vec[i].value != vec[j].value {
			return vec[i].value > vec[j].value
		}
		return vec[i].key < vec[j].key
	})

	var genreSeeds []string
	for _, vals := range vec[:5] {
		genreSeeds = append(genreSeeds, vals.key)
	}

	return artistSeeds, genreSeeds
}

func hitSeedsApi(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

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
	return s
}

type KV struct {
	key   string
	value int
}

func mapToSlice(in map[string]int) []KV {
	vec := make([]KV, len(in))
	i := 0
	for k, v := range in {
		vec[i].key = k
		vec[i].value = v
		i++
	}
	return vec
}
