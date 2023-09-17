package main
import (
 "strings"
 "github.com/antchfx/htmlquery"
)
func main() {
 doc, err := htmlquery.LoadURL("{URL}")
println(doc)
 if err != nil {
 panic(err)
 }
 TextNodes := htmlquery.Find(doc,`//{XPATH_HERE}//text()`)
 if err != nil {
 panic(err)
 }
 
}
