package auth

import (
	"encoding/json"
	"os"

	"github.com/sknain/astracore/broker-go/internal/fyers"
)

func NewFyersClient() (*fyers.Client, error) {

	data, err := LoadToken()
	if err != nil {
		return nil, err
	}

	var token TokenResponse

	if err := json.Unmarshal(data, &token); err != nil {
		return nil, err
	}

	return fyers.NewClient(
		os.Getenv("FYERS_APP_ID"),
		token.AccessToken,
	), nil
}
