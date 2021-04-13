package main

import (
	"fmt"
	"log"
	"net/http"
)

const https = "https://euw1.api.riotgames.com"
const path = "/lol/summoner/v1/summoners/by-name/mouton211"
const APIkey = "?api_key=RGAPI-d2a846d6-a608-453a-9c5d-3ecbc0a1c3d1" //valable 24h

func api(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}

func main() {
	url := https + path + APIkey
	fmt.Println(url)
	fmt.Println()

	api(url)
}
