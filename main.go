package main

import (
	"wat-r-u-doing-bot/app/controllers"
	"log"
	"os"
	"wat-r-u-doing-bot/config"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.OpenGORMConnection()
}

func main() {
	log.Printf("Hello")
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
