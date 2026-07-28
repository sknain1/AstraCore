package auth

import "net/http"

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WebSocket module coming soon"))
}
