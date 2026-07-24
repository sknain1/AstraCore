package auth

import "net/http"

func FundsHandler(w http.ResponseWriter, r *http.Request) {

	funds, err := GetFunds()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(funds)
}
