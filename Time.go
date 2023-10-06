package main

import (
    "fmt"
    "time")

func main() {
    currentTime := time.Now()
    fmt.Println("Current Time is :", currentTime)
    Tomm := currentTime.Add(24 * time.Hour)
    pastTime := currentTime.Add(-1 * time.Hour)
    fmt.Println("Future Time:", Tomm)
    fmt.Println("Past Time:", pastTime)
    Sleephrs := 10* time.Hour
    time.Sleep(Sleephrs)
}
