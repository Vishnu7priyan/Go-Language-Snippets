
package main
import (  "fmt"  "strconv"  "github.com/PuerkitoBio/goquery")
func main() {  
  doc, err := goquery.NewDocument("URL")
  if err != nil {   
    panic(err)  }  
  doc.Find(`css selector here `).   
  Each(func(i int, e *goquery.Selection) {     
    a,_ = e.Attr("change here")     
    b, _ := e.Attr("change here")     
    c, err = strconv.ParseFloat(b, 64)     
    if err != nil {        
      println("Failed to parse ")      }    
    fmt.Printf("%s ($%0.2f)\n", a, c)    })}
