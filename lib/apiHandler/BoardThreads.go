package apiHandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// a recreation of the json data from the api request.
// []Page referes to the page of the post
// []Threads refers to the threads on each page
// attributes No ... Replies are from a []Thread
type ReqBoardThreads []struct {
	Page    int `json:"page"`
	Threads []struct {
		No           int `json:"no"`
		LastModified int `json:"last_modified"`
		Replies      int `json:"replies"`
	} `json:"threads"`
}

// Simple wrapper for creating a http req, and returning
// a struct with a board listing.
// 'b' is the board to fetch data from
func GetBoardThreads(b string) ReqBoardThreads {

	// reqest data from api with 't' being the board
	req, err := http.Get("https://a.4cdn.org/" + b + "/threads.json")
	if err != nil {
		panic(err)
	}

	// defer closing
	defer req.Body.Close()

	// read data to []byte
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	// create a struct with the data
	data := ReqBoardThreads{}

	// interpert + return
	json.Unmarshal([]byte(body), &data)
	return data
}
