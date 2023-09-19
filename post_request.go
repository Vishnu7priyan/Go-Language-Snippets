package main
import(  "net/http"  
        "net/url")
func main() {  
   data := url.Values{}  
   data.Set("s", "Golang")  
   response, err :=http.PostForm("URL", data) 
}
