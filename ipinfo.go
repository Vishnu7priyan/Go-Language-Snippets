package mypackage
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)
func Ipinfo(ip string) {
    response, err := http.Get("https://ipinfo.io/" + ip + "/json")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer response.Body.Close()
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(string(data))
}
