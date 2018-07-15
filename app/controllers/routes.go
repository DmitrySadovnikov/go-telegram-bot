package controllers

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type webServer struct {
	Server http.Server
}

func NewServer() *webServer {
	s := &webServer{Server: http.Server{Addr: ":" + os.Getenv("PORT")}}
	router := Router()
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	s.Server.Handler = loggedRouter
	return s
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello\n")) })
	r.HandleFunc("/messages", Messages).Methods("POST")
	return r
}

func (s *webServer) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	<-irqSig

	log.Printf("Stopping...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.Server.Shutdown(ctx)
	if err != nil {
		log.Printf("Error while stopping: %v", err)
	}
}
