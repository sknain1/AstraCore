package auth

import (
	"encoding/json"
	"net/http"
)

func PlaceOrderHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order PlaceOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if order.Symbol == "" {
		http.Error(w, "symbol is required", http.StatusBadRequest)
		return
	}

	if order.Qty <= 0 {
		http.Error(w, "qty must be greater than zero", http.StatusBadRequest)
		return
	}

	// ------------------------------
	// DRY RUN MODE
	// ------------------------------
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]any{
		"status":  "dry-run",
		"message": "Order validated successfully. No order sent to Fyers.",
		"order":   order,
	})

	return

	// Live mode (अभी disabled है)
	/*
	resp, err := PlaceOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	*/
}
