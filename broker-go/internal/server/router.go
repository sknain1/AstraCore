package server

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/version", VersionHandler)
	http.HandleFunc("/ping", PingHandler)
}
