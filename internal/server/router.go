package server

import (
	"net/http"

	"github.com/sknain/astracore/broker-go/internal/auth"
	"github.com/sknain/astracore/broker-go/internal/order"
)

func RegisterRoutes() http.Handler {

	mux := http.NewServeMux()

	// Health APIs
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/version", VersionHandler)
	mux.HandleFunc("/ping", PingHandler)

	// Authentication APIs
	mux.HandleFunc("/auth/login", auth.LoginHandler)
	mux.HandleFunc("/auth/callback", auth.CallbackHandler)

	// Account APIs
	mux.HandleFunc("/profile", auth.ProfileHandler)
	mux.HandleFunc("/funds", auth.FundsHandler)
	mux.HandleFunc("/holdings", auth.HoldingsHandler)
	mux.HandleFunc("/positions", auth.PositionsHandler)
	mux.HandleFunc("/orders", auth.OrdersHandler)
	mux.HandleFunc("/ltp", auth.LTPHandler)
	mux.HandleFunc("/history", auth.HistoryHandler)

	// WebSocket APIs
	mux.HandleFunc("/ws/start", auth.StartWSHandler)
	mux.HandleFunc("/ws/stop", auth.StopWSHandler)
	mux.HandleFunc("/ws/subscribe", auth.SubscribeWSHandler)
	mux.HandleFunc("/ws/status", auth.WSStatusHandler)
	mux.HandleFunc("/ws/ticks", auth.LatestTickHandler)

	// Order Engine APIs
	mux.HandleFunc("/orders/place", order.PlaceOrderHandler)
	mux.HandleFunc("/orders/modify", order.ModifyOrderHandler)
	mux.HandleFunc("/orders/cancel", order.CancelOrderHandler)
	mux.HandleFunc("/orders/list", order.ListOrdersHandler)
	mux.HandleFunc("/orders/status", order.GetOrderHandler)

	// Middlewares
	handler := LoggingMiddleware(mux)
	handler = RecoveryMiddleware(handler)

	return handler
}
