package main

import (
    "./app/controllers"
    "log"
    "os"
)

func main() {
    server := controllers.NewServer()

    go func() {
        log.Printf("Server running on port " + os.Getenv("PORT") + "...")
        err := server.Server.ListenAndServe()
        if err != nil {
            log.Printf("Error: %v", err)
        }
    }()

    server.WaitShutdown()
    log.Printf("Server shuted down")
}
