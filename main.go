package main

import (
	"fmt"
	"log"
	"net/http"
)

func api(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}

func main() {
	url := "https://127.0.0.1:2999/liveclientdata/activeplayer"
	api(url)
}
