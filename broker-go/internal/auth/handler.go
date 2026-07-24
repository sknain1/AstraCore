package auth

import "net/http"

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	profile, err := GetProfile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(profile)
}
