package user

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func createTopUserRequests(url, term, limit, AccessToken string) string {

	Authorization := "Bearer " + AccessToken

	req, err := http.NewRequest("GET", url, nil)

	if term == "short" {
		term = "short_term"
	} else {
		term = "medium_term"
	}

	q := req.URL.Query()
	q.Add("time_range", term)
	q.Add("limit", limit)
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

	return s
}
