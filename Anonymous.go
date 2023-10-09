package main

import "fmt"

func main(){
  add := func(x, y int) int {
    return x + y
  }
  result := add(3, 5)
  prin := func(){
    fmt.Println("Hii")
    }
  prin()
}
