package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sknain/astracore/broker-go/internal/config"
	"github.com/sknain/astracore/broker-go/internal/server"
)

func main() {

	cfg, err := config.Load("../configs/app.yaml")
	if err != nil {
		log.Fatal(err)
	}

	handler := server.RegisterRoutes()

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

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
