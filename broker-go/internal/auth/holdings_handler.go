package auth

import "net/http"

func HoldingsHandler(w http.ResponseWriter, r *http.Request) {

	holdings, err := GetHoldings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(holdings)
}
