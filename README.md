This repository consist of simple Go Language snippets

1. Anonymous Add and Print(https://github.com/Vishnu7priyan/Go-Lang/blob/4c9d6bca2e3bcaededb41bdbec5aa02fa34ef9ea/Anonymous.go)

``` Go
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
```
