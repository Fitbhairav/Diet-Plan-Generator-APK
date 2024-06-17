package main

import (
    "log"
    "net/http"

    "github.com/joho/godotenv"
    "diet-plan-generator/database"
    "diet-plan-generator/routers"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found")
    }

    database.InitDB()
    routers.InitializeRoutes()

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
