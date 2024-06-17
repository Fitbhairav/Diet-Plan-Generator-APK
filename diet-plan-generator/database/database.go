package database

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitDB() {
    clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_DSN"))
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Failed to ping database:", err)
    }

    DB = client.Database("diteplan") 
    log.Println("Connected to MongoDB!")
}
