package auth

import (
	"encoding/json"
	"net/http"

	"github.com/sknain/astracore/broker-go/internal/ws"
)

type SubscribeRequest struct {
	Symbols []string `json:"symbols"`
}

func StartWSHandler(w http.ResponseWriter, r *http.Request) {
	if err := ws.GetManager().Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "WebSocket started",
	})
}

func StopWSHandler(w http.ResponseWriter, r *http.Request) {
	ws.GetManager().Stop()

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "WebSocket stopped",
	})
}

func SubscribeWSHandler(w http.ResponseWriter, r *http.Request) {
	var req SubscribeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ws.GetManager().Subscribe(req.Symbols, "SymbolUpdate"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"symbols": req.Symbols,
	})
}

func WSStatusHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"connected": ws.GetManager().Status(),
	})
}

func LatestTickHandler(w http.ResponseWriter, r *http.Request) {

	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		http.Error(w, "missing symbol", http.StatusBadRequest)
		return
	}

	tick, ok := ws.GetTick(symbol)
	if !ok {
		http.Error(w, "tick not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tick)
}
