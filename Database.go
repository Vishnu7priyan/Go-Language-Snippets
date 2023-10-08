
package main

import (
    "database/sql"
    "fmt"
    "github.com/mattn/go-sqlite3"
    //Install with go get <URL>
)

func main() {
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        fmt.Println("Error :", err)
    }
    //Table creation
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

    //Updation
    updateSQL := "UPDATE users SET email = ? WHERE username = ?"
    _, err = db.Exec(updateSQL, "test@example.com", "vishnu")
    if err != nil {
        fmt.Println("Error updating data:", err)
        return
    }
    //Deletion
    deleteSQL := "DELETE FROM users WHERE username = ?"
    _, err = db.Exec(deleteSQL, "vishnu")
    if err != nil {
        fmt.Println("Error deleting data:", err)
        return
    }

    
}
