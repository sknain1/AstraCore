package auth

import "net/http"

func PositionsHandler(w http.ResponseWriter, r *http.Request) {

	positions, err := GetPositions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(positions)
}
