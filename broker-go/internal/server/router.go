package server

import "net/http"

func RegisterRoutes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/version", VersionHandler)
	mux.HandleFunc("/ping", PingHandler)

	handler := LoggingMiddleware(mux)
	handler = RecoveryMiddleware(handler)

	return handler
}
