package auth

func GetHoldings() ([]byte, error) {

	client, err := NewFyersClient()
	if err != nil {
		return nil, err
	}

	return client.Get("holdings")
}
