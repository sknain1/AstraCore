package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/sknain/astracore/broker-go/internal/auth"
	"github.com/sknain/astracore/broker-go/internal/broker"
	"github.com/sknain/astracore/broker-go/internal/config"
	"github.com/sknain/astracore/broker-go/internal/fyers"
	"github.com/sknain/astracore/broker-go/internal/server"
	"github.com/sknain/astracore/broker-go/internal/ws"
)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Warning: .env file not found")
	}

	cfg, err := config.Load("../configs/app.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Register Event Subscribers
	ws.RegisterSubscribers()

	// Load Token Once
	token, err := auth.LoadToken()
	if err != nil {

		log.Printf("Broker/WebSocket disabled: %v", err)

	} else {

		// -----------------------------
		// Register FYERS Broker
		// -----------------------------
		fyersClient := fyers.NewClient(
			os.Getenv("FYERS_APP_ID"),
			token.AccessToken,
		)

		adapter := fyers.NewAdapter(fyersClient)

		broker.Register(adapter)

		log.Println("FYERS broker registered")

		// -----------------------------
		// Initialize WebSocket
		// -----------------------------
		wsClient := ws.NewClient(
			os.Getenv("FYERS_APP_ID"),
			token.AccessToken,
		)

		ws.GetManager().SetClient(wsClient)

		log.Println("WebSocket client initialized")
	}

	handler := server.RegisterRoutes()

	addr := fmt.Sprintf("%s:%d",
		cfg.Server.Host,
		cfg.Server.Port,
	)

	log.Printf("%s v%s (%s) started on %s",
		cfg.App.Name,
		cfg.App.Version,
		cfg.App.Env,
		addr,
	)

	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
