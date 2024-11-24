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
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	rss, _ := json.Marshal(data)
	fmt.Fprint(w, string(rss))
}
