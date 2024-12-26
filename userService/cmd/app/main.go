package main

import (
	"log"
	"net/http"
	"time"
	"userService/internal/configs"
	"userService/internal/handlers"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlers.MainHandler)

	cfg := configs.Get()

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      m,
		ReadTimeout:  cfg.Timeout * time.Second,
		WriteTimeout: cfg.Timeout * time.Second,
	}

	log.Printf("Listening %s...", cfg.Port)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
