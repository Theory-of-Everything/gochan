package apiHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// a recreation of a thread from an api reqest
type ReqThread struct {
	Posts []struct {
		No          int    `json:"no"`
		Sticky      int    `json:"sticky"`
		Closed      int    `json:"closed"`
		Now         string `json:"now"`
		Name        string `json:"name"`
		Sub         string `json:"sub"`
		Com         string `json:"com"`
		Filename    string `json:"filename"`
		Ext         string `json:"ext"`
		W           int    `json:"w"`
		H           int    `json:"h"`
		TnW         int    `json:"tn_w"`
		TnH         int    `json:"tn_h"`
		Tim         int64  `json:"tim"`
		Time        int    `json:"time"`
		Md5         string `json:"md5"`
		Fsize       int    `json:"fsize"`
		Resto       int    `json:"resto"`
		Capcode     string `json:"capcode"`
		SemanticURL string `json:"semantic_url,omitempty"`
		Replies     int    `json:"replies,omitempty"`
		Images      int    `json:"images,omitempty"`
		UniqueIps   int    `json:"unique_ips,omitempty"`
	} `json:"posts"`
}

// Simple wrapper for creating a http req, and returning
// a struct with a thread listing.
// 'b' is the board to fetch data from
// 't' is the id of the thread
func GetThread(b string, t int) ReqThread {

	// reqest the content of given thread
	req, err := http.Get("https://a.4cdn.org/" + b + "/thread/" + fmt.Sprint(t) + ".json")
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

	// create data struct
	data := ReqThread{}

	// write + return
	json.Unmarshal([]byte(body), &data)
	return data
}
