package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

func main() {
    response, err := http.Get("https://web-check.as93.net/.netlify/functions/dns?url=https://ikea.com")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    if response.StatusCode == http.StatusOK {
        fmt.Println("GET request successful")

        responseBody, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatal(err)
        }

        var jsonData interface{}

        if err := json.Unmarshal(responseBody, &jsonData); err != nil {
            log.Fatal(err)
        }

      
}
