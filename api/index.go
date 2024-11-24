package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Feeds []struct {
		Name string `json:"name"`
		Uid  int    `json:"uid"`
		Url  string `json:"url"`
		Desc string `json:"desc"`
	} `json:"feeds"`
}

var data Data

func init() {
	// open a file
	f, err := os.Open("../data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read the file
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal the file
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	rss, _ := json.Marshal(data)
	fmt.Fprint(w, string(rss))
}
