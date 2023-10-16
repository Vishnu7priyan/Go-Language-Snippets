package main

import (
    "fmt"
    "net/http"
    ")

func main(){
  url := "URL_HERE"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("GET Response:")
	fmt.Println(string(body))

  jsonData := `{"title": "Sample Post", "body": "This is a sample", "userId": 1}`
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("POST Response Status:", resp.Status)
  }
