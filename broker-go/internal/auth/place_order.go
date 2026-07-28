package auth

import (
	"encoding/json"
	"fmt"
)

func PlaceOrder(order PlaceOrderRequest) ([]byte, error) {

	client, err := NewFyersClient()
	if err != nil {
		return nil, err
	}

	payload, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		return nil, err
	}

	fmt.Println("===================================")
	fmt.Println("FYERS ORDER REQUEST")
	fmt.Println(string(payload))
	fmt.Println("===================================")

	resp, err := client.Post("orders/sync", payload)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
