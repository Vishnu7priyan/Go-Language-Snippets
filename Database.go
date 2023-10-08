
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
    createTable := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT,
            email TEXT
        )
    `
    _, err = db.Exec(createTable)
    if err != nil {
        fmt.Println("Error :", err)
    }
    
    updateSQL := "UPDATE users SET email = ? WHERE username = ?"
    _, err = db.Exec(updateSQL, "test@example.com", "vishnu")
    if err != nil {
        fmt.Println("Error updating data:", err)
        return
    }
    

    
}
