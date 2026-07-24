package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetHoldings() ([]byte, error) {

	data, err := LoadToken()
	if err != nil {
		return nil, err
	}

	var token TokenResponse

	if err := json.Unmarshal(data, &token); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodGet,
		"https://api-t1.fyers.in/api/v3/holdings",
		nil,
	)
	if err != nil {
		return nil, err
	}

	appID := os.Getenv("FYERS_APP_ID")

	req.Header.Set(
		"Authorization",
		appID+":"+token.AccessToken,
	)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}
