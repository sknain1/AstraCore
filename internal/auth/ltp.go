package auth

func GetLTP(symbol string) ([]byte, error) {

	client, err := NewFyersClient()
	if err != nil {
		return nil, err
	}

	return client.Get("data/quotes?symbols=" + symbol)
}
