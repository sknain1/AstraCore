package auth

func GetProfile() ([]byte, error) {

	client, err := NewFyersClient()
	if err != nil {
		return nil, err
	}

	return client.Get("profile")
}
