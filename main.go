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

const APIkey = "?api_key=RGAPI-549361bd-d625-4629-b7c2-dc246de49620" //valable 24h
const accountId = "D8lq_rQ9lXxYdl867SpZMRo6UqH7fAJ3hmaQSJ8sbqlPrV8"
const puuid = "Y6dqJmXGwwagdqXz27gSpv5Mf_J5xZ2owmIcpek6LLvJnSJ0nPwCY984LzDWIomm1omlLIShODcenw"

const tagLine = "EUW"
const gameName = "mouton211"

// 20 requests every 1 seconds(s) & 100 requests every 2 minutes(s)
type playerData struct {
    Level string
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

/*func apiclient(urlclient string) {

	response, err := http.Get(urlclient)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response, string(body))

}*/
func main() {
    for{
	urlprofil := https + "/lol/summoner/v4/summoners/by-name/" + gameName + APIkey
	//urlclient := "https://127.0.0.1:2999/liveclientdata/allgamedata"

	apiprofil(urlprofil)
	fmt.Println()
	//apiclient(urlclient)

	fmt.Println("###########################################################")
	fmt.Println()

	fmt.Println("Nom du joueur:", gameName, "  Niveaux", "  Serveur:", tagLine)
	fmt.Println()

	fmt.Println("###########################################################")
	
	fmt.Println()
	time.Sleep(100*time.Millisecond)
    }
}
