package server

import (
	"net/http"

	"github.com/sknain/astracore/broker-go/internal/auth"
)

func RegisterRoutes() http.Handler {

	mux := http.NewServeMux()

	// Health APIs
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/version", VersionHandler)
	mux.HandleFunc("/ping", PingHandler)

	// Authentication APIs
	mux.HandleFunc("/auth/login", auth.LoginHandler)
	mux.HandleFunc("/auth/callback", auth.CallbackHandler)

	// Profile API
	mux.HandleFunc("/profile", auth.ProfileHandler)
        mux.HandleFunc("/funds", auth.FundsHandler)
	// Middlewares
	handler := LoggingMiddleware(mux)
	handler = RecoveryMiddleware(handler)

	return handler
}
