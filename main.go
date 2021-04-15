package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const https = "https://euw1.api.riotgames.com"
const APIkey = "?api_key=RGAPI-6c5f3688-5786-4b0b-b6cd-b319482d0e96" //valable 24h
const accountId = "D8lq_rQ9lXxYdl867SpZMRo6UqH7fAJ3hmaQSJ8sbqlPrV8"
const puuid = "Y6dqJmXGwwagdqXz27gSpv5Mf_J5xZ2owmIcpek6LLvJnSJ0nPwCY984LzDWIomm1omlLIShODcenw"
const tagLine = "EUW"
const gameName = "mouton211"

// 20 requests every 1 seconds(s) & 100 requests every 2 minutes(s)
type playerData struct {
	SummonerLevel int16
	Name          string
}

func apiProfil(urlprofil string) playerData {
	response, err := http.Get(urlprofil)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	apiData := playerData{}
	err = json.Unmarshal(body, &apiData)
	if err != nil {
		log.Fatal(err)
	}

	return apiData
}

func apiStat(urlstat string) {
	response, err := http.Get(urlstat)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response, string(body))
}

func main() {
	for {
		urlprofil := https + "/lol/summoner/v4/summoners/by-name/" + gameName + APIkey
		playerInfo := apiProfil(urlprofil)

		urlstat := "https://127.0.0.1:2999/liveclientdata/allgamedata"
		apiStat(urlstat)

		fmt.Println("###########################################################")
		fmt.Println("Nom du joueur:", playerInfo.Name, "  Niveau:", playerInfo.SummonerLevel, "  Serveur:", tagLine)
		fmt.Println("###########################################################")
		fmt.Println()
		time.Sleep(100 * time.Millisecond)
	}
}
