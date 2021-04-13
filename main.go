package main

import (
	"fmt"
	"log"
	"net/http"
)

const https = "https://"
const path = "euw1.api.riotgames.com/lol/platform/v3/champion-rotations?"
const APIkey = "api_key=RGAPI-d2a846d6-a608-453a-9c5d-3ecbc0a1c3d1" //valable 24h

func api(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}

func main() {
	url := https + path + APIkey
	api(url)
}
