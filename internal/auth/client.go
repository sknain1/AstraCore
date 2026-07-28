package auth

import (
	"encoding/json"
	"fmt"
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

	fmt.Println("APP ID =", os.Getenv("FYERS_APP_ID"))
	fmt.Println("TOKEN PREFIX =", token.AccessToken[:20])

	return fyers.NewClient(
		os.Getenv("FYERS_APP_ID"),
		token.AccessToken,
	), nil
}
