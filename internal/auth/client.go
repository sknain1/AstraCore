package auth

import (
	"fmt"
	"os"

	"github.com/sknain/astracore/broker-go/internal/fyers"
)

func NewFyersClient() (*fyers.Client, error) {

	token, err := LoadToken()
	if err != nil {
		return nil, err
	}

	fmt.Println("APP ID =", os.Getenv("FYERS_APP_ID"))

	if len(token.AccessToken) >= 20 {
		fmt.Println("TOKEN PREFIX =", token.AccessToken[:20])
	} else {
		fmt.Println("TOKEN PREFIX =", token.AccessToken)
	}

	return fyers.NewClient(
		os.Getenv("FYERS_APP_ID"),
		token.AccessToken,
	), nil
}
