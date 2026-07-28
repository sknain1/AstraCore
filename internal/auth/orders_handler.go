package auth

import "net/http"

func OrdersHandler(w http.ResponseWriter, r *http.Request) {

	orders, err := GetOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(orders)
}
