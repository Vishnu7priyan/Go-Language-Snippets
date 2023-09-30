package main

import(
  "fmt"
  )
func main(){
  colors := map[string] string{
    "Blue":"#f0f8ff",
    "Coral":"#ff7f50",
  }
  for key,value := range colors{
    fmt.Println("%s and %s " , key,value}
}
delete(colors,"Blue")
