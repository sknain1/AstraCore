package auth

import (
	"os"
)

const tokenFile = "token.json"

func SaveToken(token []byte) error {
	return os.WriteFile(tokenFile, token, 0600)
}

func LoadToken() ([]byte, error) {
	return os.ReadFile(tokenFile)
}
