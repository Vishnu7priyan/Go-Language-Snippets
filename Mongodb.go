package main

import (
    "context"
    "fmt"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        fmt.Println(err)
        return
    }

    db := client.Database("mydb")
    collection := db.Collection("mycollection")

    document := bson.D{{"name", "John Doe"}, {"age", 30}}

    
}
