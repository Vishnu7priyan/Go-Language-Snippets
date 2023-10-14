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
    defer resp.Body.Close()
     data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    stringBody := string(data)

    re := regexp.MustCompile(`<a.*href\s*=\s*["'](https?://[^"']+)["'].*>`)
    links := re.FindAllStringSubmatch(stringBody, -1)

    fmt.Println(len(links))
    for _, link := range links {
        fmt.Println(link[1])
    }
    
   
}
