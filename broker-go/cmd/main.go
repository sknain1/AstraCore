package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sknain/astracore/broker-go/internal/server"
)

func main() {

	handler := server.RegisterRoutes()

	log.Println("Broker Service started on :8080")

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
