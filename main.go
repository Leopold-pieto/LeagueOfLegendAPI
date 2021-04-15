package main

import (
	"crypto/tls"
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
type gameData struct {
	CurrentGold float64
	GameTime    float64
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

	apiData := playerData{}
	err = json.Unmarshal(body, &apiData)
	if err != nil {
		log.Fatal(err)
	}

	return apiData
}

func apiGold(urlgold string) gameData {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	response, err := http.Get("https://127.0.0.1:2999/liveclientdata/activeplayer")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	goldData := gameData{}
	err = json.Unmarshal(body, &goldData)
	if err != nil {
		log.Fatal(err)
	}
	return goldData
}
func apitime(urltime string) gameData {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	response, err := http.Get("https://127.0.0.1:2999/liveclientdata/gamestats")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	timeData := gameData{}
	err = json.Unmarshal(body, &timeData)
	if err != nil {
		log.Fatal(err)
	}
	return timeData

}

func main() {
	for {
		urlprofil := https + "/lol/summoner/v4/summoners/by-name/" + gameName + APIkey
		playerInfo := apiProfil(urlprofil)

		urlgold := "https://127.0.0.1:2999/liveclientdata/activeplayer"
		playerGold := apiGold(urlgold)

		urltime := "https://127.0.0.1:2999/liveclientdata/gamestats"
		gameTime := apitime(urltime)

		fmt.Println("###########################################################")
		fmt.Println("Nom du joueur:", playerInfo.Name, "  Niveau:", playerInfo.SummonerLevel, "  Serveur:", tagLine)
		fmt.Println("temps de jeu en seconde:", gameTime.GameTime, "Gold gagné:", playerGold.CurrentGold)
		fmt.Println("###########################################################")
		fmt.Println()
		time.Sleep(100 * time.Millisecond)
	}
}
