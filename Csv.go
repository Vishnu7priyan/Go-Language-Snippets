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

