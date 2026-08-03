package fyers

import (
	"encoding/json"

	"github.com/sknain/astracore/broker-go/internal/broker"
)

type Adapter struct {
	client *Client
}

func NewAdapter(client *Client) *Adapter {
	return &Adapter{
		client: client,
	}
}

func (a *Adapter) PlaceOrder(req broker.PlaceOrderRequest) (string, error) {

	side := 1
	if req.Side == broker.Sell {
		side = -1
	}

	resp, err := a.client.PlaceOrder(PlaceOrderRequest{
		Symbol:       req.Symbol,
		Qty:          req.Qty,
		Type:         1,
		Side:         side,
		ProductType:  "INTRADAY",
		LimitPrice:   req.Price,
		StopPrice:    0,
		DisclosedQty: 0,
		Validity:     "DAY",
		OfflineOrder: false,
	})

	if err != nil {
		return "", err
	}

	var result map[string]any

	if err := json.Unmarshal(resp, &result); err != nil {
		return "", err
	}

	id, _ := result["id"].(string)

	return id, nil
}

func (a *Adapter) ModifyOrder(orderID string, qty int, price float64) error {

	_, err := a.client.ModifyOrder(ModifyOrderRequest{
		ID:         orderID,
		Qty:        qty,
		LimitPrice: price,
	})

	return err
}

func (a *Adapter) CancelOrder(orderID string) error {

	_, err := a.client.CancelOrder(orderID)

	return err
}

func (a *Adapter) GetOrder(orderID string) ([]byte, error) {

	return a.client.OrderStatus(orderID)
}
