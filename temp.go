package main

import (
    "fmt"
    "log"
    "net/http"
     "encoding/json"
    "io/ioutil"
   
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

        var jsonData interface{}

        if err := json.Unmarshal(responseBody, &jsonData); err != nil {
            log.Fatal(err)
        }

        formattedJSON, err := json.MarshalIndent(jsonData, "", "    ")
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("JSON Response:")
        fmt.Println(string(formattedJSON))
    } else {
        log.Printf("GET request failed with status: %s\n", response.Status)
    }
}
