package order

import (
	"encoding/json"
	"net/http"
)

var service = NewService()

type PlaceOrderRequest struct {
	Symbol string  `json:"symbol"`
	Side   Side    `json:"side"`
	Qty    int     `json:"qty"`
	Price  float64 `json:"price"`
}

type ModifyOrderRequest struct {
	ID    string  `json:"id"`
	Qty   int     `json:"qty"`
	Price float64 `json:"price"`
}

type CancelOrderRequest struct {
	ID string `json:"id"`
}

func PlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req PlaceOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := service.Place(req.Symbol, req.Side, req.Qty, req.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func ModifyOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req ModifyOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := service.Modify(req.ID, req.Qty, req.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req CancelOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := service.Cancel(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(service.All())
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	order, ok := service.Get(id)
	if !ok {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}
