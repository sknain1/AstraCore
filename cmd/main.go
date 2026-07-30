package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sknain/astracore/broker-go/internal/auth"
	"github.com/sknain/astracore/broker-go/internal/broker"
	"github.com/sknain/astracore/broker-go/internal/config"
	"github.com/sknain/astracore/broker-go/internal/fyers"
	"github.com/sknain/astracore/broker-go/internal/instruments"
	"github.com/sknain/astracore/broker-go/internal/market"
	"github.com/sknain/astracore/broker-go/internal/paper"
	"github.com/sknain/astracore/broker-go/internal/server"
	"github.com/sknain/astracore/broker-go/internal/ws"
)

func main() {

	// Load .env
	if err := config.LoadEnv(); err != nil {
		log.Println(err)
	}

	// Load Config
	cfg, err := config.Load("../configs/app.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Load Instruments
	if err := instruments.Load(); err != nil {
		log.Println(err)
	} else {
		log.Printf("Loaded %d instruments", len(instruments.All()))
	}

	// Register Event Subscribers
	ws.RegisterSubscribers()
	market.RegisterSubscribers()

	// Load Token
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

		fyersAdapter := fyers.NewAdapter(fyersClient)
		broker.Register("fyers", fyersAdapter)

		// -----------------------------
		// Register Paper Broker
		// -----------------------------
		paperAdapter := paper.NewAdapter()
		broker.Register("paper", paperAdapter)

		// -----------------------------
		// Select Active Broker
		// -----------------------------
		brokerName := os.Getenv("BROKER")
		if brokerName == "" {
			brokerName = "paper"
		}

		if err := broker.Use(brokerName); err != nil {
			log.Fatal(err)
		}

		log.Printf("Active Broker: %s", brokerName)

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

	addr := fmt.Sprintf(
		"%s:%d",
		cfg.Server.Host,
		cfg.Server.Port,
	)

	log.Printf(
		"%s v%s (%s) started on %s",
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
