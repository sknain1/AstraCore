package auth

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const tokenURL = "https://api-t1.fyers.in/api/v3/validate-authcode"

type tokenRequest struct {
	AppIDHash string `json:"appIdHash"`
	Code      string `json:"code"`
	GrantType string `json:"grant_type"`
}

func ExchangeToken(authCode string) ([]byte, error) {

	appID := os.Getenv("FYERS_APP_ID")
	secret := os.Getenv("FYERS_SECRET_KEY")

	hash := sha256.Sum256([]byte(appID + ":" + secret))
	appIDHash := hex.EncodeToString(hash[:])

	reqBody := tokenRequest{
		AppIDHash: appIDHash,
		Code:      authCode,
		GrantType: "authorization_code",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		tokenURL,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", data)
	}

	return data, nil
}
