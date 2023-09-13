//Default commands
package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("cd", "/tmp")
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    }
}



