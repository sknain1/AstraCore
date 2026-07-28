package auth

import (
	"net/http"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {

	authCode := r.URL.Query().Get("auth_code")

	if authCode == "" {
		http.Error(w, "missing auth_code", http.StatusBadRequest)
		return
	}

	token, err := ExchangeToken(authCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save Access Token
	if err := SaveToken(token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Access Token saved successfully."))
}
