package auth

import (
	"encoding/json"
	"os"
)

const tokenFile = "token.json"

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
	Status       string `json:"s"`
}

func SaveToken(data []byte) error {
	return os.WriteFile(tokenFile, data, 0600)
}

func LoadToken() (*Token, error) {
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		return nil, err
	}

	var token Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, err
	}

	return &token, nil
}
