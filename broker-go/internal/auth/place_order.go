package auth

import "encoding/json"

func PlaceOrder(order PlaceOrderRequest) ([]byte, error) {

	client, err := NewFyersClient()
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	return client.Post("orders", payload)
}
