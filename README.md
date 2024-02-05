This repository consist of simple Go Language snippets

1. Anonymous Add and Print 
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
2.Ansi color
```go
package main

import (
        "fmt"
)

func main() {
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
        
    fmt.Println(yellow + "This is yellow")
    fmt.Println(green + "This is green")
    fmt.Println(red + "This is red")
    /*
    More colors
    Black: \033[30m
    Blue: \033[34m
    Magenta: \033[35m
    Cyan: \033[36m
    White: \033[37m
    */
}
```
