package main

import (  "fmt"  
"io/ioutil"  
"net/http"  
"strings")

func main() {  
   resp, err := http.Get("URL")  
   if err != nil {    
      panic(err)  }  
   data, err := ioutil.ReadAll(resp.Body) 
   if err != nil {   
      panic(err)  }
//some more code
}
