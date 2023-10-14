//html-tool is developed by tomnomnom 

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "os/exec"
    execute "github.com/alexellis/go-execute/pkg/v1"
)

func GetAllUrls(user_input string) {
    url := "https://"+user_input
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close() // Close the response body to avoid resource leaks.

   
}
