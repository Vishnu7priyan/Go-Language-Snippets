package main

import (
    "flag"
    "fmt"
)

func main() {
    var yourname string
    var age int

    flag.StringVar(&yourname, "name", "", "Your name")
    flag.IntVar(&age, "age", 0, "Your age")

    flag.Parse()

    
}
