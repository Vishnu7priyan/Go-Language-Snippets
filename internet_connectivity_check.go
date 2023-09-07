package main
import (
        "fmt"
        "net"
        "os"
)
func main(){
        host := "google.com"
        port := "80"
        conn,err := net.Dial("tcp", host+":"+port)
        if err != nil {
                fmt.Println("No internet connection")
                os.Exit(1)
        }
        defer conn.Close()

        fmt.Println("Internet connection is available")
}
