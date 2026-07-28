package auth

func GetPositions() ([]byte, error) {

	client, err := NewFyersClient()
	if err != nil {
		return nil, err
	}

	return client.Get("positions")
}
