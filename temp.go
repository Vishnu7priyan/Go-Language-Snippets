package main

import (
    "fmt"
    "log"
    "net/http"
   
)

func main() {
    response, err := http.Get("URL")
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

       
}
