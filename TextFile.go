package main

import (
    "fmt"
    "os"
)

func main() {
    filePath := "textfile.txt"
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening the file:", err)
        return
    }
    defer file.Close() 

}

