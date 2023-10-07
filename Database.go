
package main

import (
    "database/sql"
    "fmt"
    "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println("Error :", err)
    }
    defer db.Close()
}
