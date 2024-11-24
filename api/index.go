package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Data struct {
	Feeds []struct {
		Name string `json:"name"`
		Uid  int    `json:"uid"`
		Url  string `json:"url"`
		Desc string `json:"desc"`
	} `json:"feeds"`
}

const (
	URL = "https://raw.githubusercontent.com/colorsakura/rss-proxy/refs/heads/rss/data.json"
)

var data Data

func init() {
	data = getData()
}

func getData() Data {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result Data
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rss, _ := json.Marshal(data)
	fmt.Fprint(w, string(rss))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// TODO: routes handle
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	mux.ServeHTTP(w, r)
}
