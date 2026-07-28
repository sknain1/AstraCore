package auth

func GetOrders() ([]byte, error) {

	client, err := NewFyersClient()
	if err != nil {
		return nil, err
	}

	return client.Get("orders")
}
