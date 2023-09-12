package main

import (
"net/http"
"os"
"fmt"
)

func main(){
r , err := http.Get("http://www.example.com/index.html")
if err != nil{
        fmt.Println("Error")
        }
bodyLen := 1300 //Change the length based on your requirement
webOut := make([]byte,bodyLen)
r.Body.Read(webOut)
}
