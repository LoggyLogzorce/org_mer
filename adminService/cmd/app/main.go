package main

import (
	"adminService/internal/configs"
	"adminService/internal/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handler.MainHandler)

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
