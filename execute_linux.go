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
//////////////////////////////////////////////
package main
import (
    "fmt"
    "os/exec"
    "runtime"
)
func main() {
    out, err := exec.Command("ls").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    output := string(out[:])
    fmt.Println(output)

  }

/////////////////////////////////////////////
//Custom shell scripts
package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("PATH_HERE")
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    }
}
