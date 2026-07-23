package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	ServiceName = "broker-go"
	Version     = "0.1.0"
)

type APIResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp,omitempty"`
}

func writeJSON(w http.ResponseWriter, data APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, APIResponse{
		Status:    "ok",
		Service:   ServiceName,
		Version:   Version,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, APIResponse{
		Status:  "ok",
		Service: ServiceName,
		Version: Version,
	})
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/ping", pingHandler)

	log.Printf("%s v%s started on :8080", ServiceName, Version)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
