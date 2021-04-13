
package main

import (
    "fmt"
    "net/http"
    
)

func main() {
    resp, err := http.Get("https://127.0.0.1:2999/liveclientdata/activeplayer")

    if err != nil {
    }

}
