package main

import (
	"github.com/joho/godotenv"
	"go-telegram-bot/app/controllers"
	"go-telegram-bot/config"
	"log"
	"os"
)

func init() {
	godotenv.Load()
	config.OpenGORMConnection()
}

func main() {
	log.Printf("Hello")
	server := controllers.NewServer()

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		log.Printf("Server running on port " + port + "...")
		err := server.Server.ListenAndServe()
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}()

	server.WaitShutdown()
	log.Printf("Server shut down")
}
