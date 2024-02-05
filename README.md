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

3.Array
```Go
//Declare
var array [5]int
//Declare with values 
array := [5]int{10, 20, 30, 40, 50}
//Capacity based on no. of values
array := [...]int{10, 20, 30, 40, 50}
//Declaring on specific index values
array := [5]int{1: 10, 2: 20}
//Change The value
array[2] = 35;
var array1 [5]string 
array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
array1 = array2

var array [4][2]int
array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
array[0][0] = 10 
array[0][1] = 20 
array[1][0] = 30

slice := []int{10, 20, 30, 40, 50} 

newSlice := slice[1:3] 

newSlice = append(newSlice, 60)

slice2 := slice[2:3:4]
```

4.Binary File
```Go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("binaryfile.bin")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}
}
```

5.CSV
```go
package main

import (
    "encoding/csv"
    "os"
)

func main() {
    // Create a CSV file for writing
    file, err := os.Create("data.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()
   writer := csv.NewWriter(file)
    defer writer.Flush()

    data := [][]string{
        {"Name", "Age", "City"},
        {"vishnu", "19", "Bash"},
    }
    for _, row := range data {
        if err := writer.Write(row); 
        err != nil {
            panic(err)
        }
    }
    
}
```

