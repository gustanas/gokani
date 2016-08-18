package gokani

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIURL is the base url of wanikani
const (
	APIURL = "https://www.wanikani.com/api"
)

// Client holds the API private key and performs the calls to the wanikani API
type Client struct {
	Key string
}

var httpClient = &http.Client{}

// Call performs the calls to the different endpoints
func Call(method, url string, v interface{}) error {
	r, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	// Parsing
	res, err := httpClient.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}
	fmt.Println("v is nil")
	return nil
}
