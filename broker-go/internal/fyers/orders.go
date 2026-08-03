package fyers

import "encoding/json"

type PlaceOrderRequest struct {
	Symbol       string  `json:"symbol"`
	Qty          int     `json:"qty"`
	Type         int     `json:"type"`
	Side         int     `json:"side"`
	ProductType  string  `json:"productType"`
	LimitPrice   float64 `json:"limitPrice"`
	StopPrice    float64 `json:"stopPrice"`
	Validity     string  `json:"validity"`
	DisclosedQty int     `json:"disclosedQty"`
	OfflineOrder bool    `json:"offlineOrder"`
	StopLoss     float64 `json:"stopLoss"`
	TakeProfit   float64 `json:"takeProfit"`
	OrderTag     string  `json:"orderTag,omitempty"`
	IsSliceOrder bool    `json:"isSliceOrder,omitempty"`
}

func (c *Client) PlaceOrder(req PlaceOrderRequest) ([]byte, error) {

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return c.Post("orders/sync", payload)
}

type ModifyOrderRequest struct {
	ID         string  `json:"id"`
	Qty        int     `json:"qty,omitempty"`
	LimitPrice float64 `json:"limitPrice,omitempty"`
	StopPrice  float64 `json:"stopPrice,omitempty"`
	Type       int     `json:"type,omitempty"`
}

func (c *Client) ModifyOrder(req ModifyOrderRequest) ([]byte, error) {

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return c.Patch("orders/sync", payload)
}

type CancelOrderRequest struct {
	ID string `json:"id"`
}

func (c *Client) CancelOrder(id string) ([]byte, error) {

	req := CancelOrderRequest{
		ID: id,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return c.Delete("orders/sync", payload)
}

func (c *Client) OrderStatus(id string) ([]byte, error) {
	return c.Get("orders?id=" + id)
}
