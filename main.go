package main

import (
        "encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"log"
	"net/http"
)

const https = "https://euw1.api.riotgames.com"
const APIkey = "?api_key=RGAPI-15c29020-b2c8-4423-9ae1-7b18e959af2a"//valable 24h
const accountId = "D8lq_rQ9lXxYdl867SpZMRo6UqH7fAJ3hmaQSJ8sbqlPrV8"
const puuid = "Y6dqJmXGwwagdqXz27gSpv5Mf_J5xZ2owmIcpek6LLvJnSJ0nPwCY984LzDWIomm1omlLIShODcenw"
const tagLine = "EUW"
const gameName = "mouton211"

// 20 requests every 1 seconds(s) & 100 requests every 2 minutes(s)
type playerData struct {
       SummonerLevel int16
       Name string
}

func apiprofil(urlprofil string) {
	response, err := http.Get(urlprofil)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	apiData:=playerData{}
        err=json.Unmarshal(body,&apiData)

	if err != nil {
		log.Fatal(err)
	}
        fmt.Println(apiData)
}

func main() {
    for{
	urlprofil := https + "/lol/summoner/v4/summoners/by-name/" + gameName + APIkey

	apiprofil(urlprofil)
	fmt.Println()
	
	fmt.Println("###########################################################")
	fmt.Println()

	fmt.Println("Nom du joueur:", "  Niveau:", "  Serveur:", tagLine)
	fmt.Println()

	fmt.Println("###########################################################")

	time.Sleep(100*time.Millisecond)
    }
}
