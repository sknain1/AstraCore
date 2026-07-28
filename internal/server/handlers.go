package server

import (
	"encoding/json"
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

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, APIResponse{
		Status:    "ok",
		Service:   ServiceName,
		Version:   Version,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, APIResponse{
		Status:  "ok",
		Service: ServiceName,
		Version: Version,
	})
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
