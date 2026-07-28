package fyers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	AppID       string
	AccessToken string
}

func NewClient(appID, accessToken string) *Client {
	return &Client{
		AppID:       appID,
		AccessToken: accessToken,
	}
}

func (c *Client) Get(endpoint string) ([]byte, error) {

	url := "https://api-t1.fyers.in/api/v3/" + endpoint

	if len(endpoint) >= 4 && endpoint[:4] == "data" {
		url = "https://api-t1.fyers.in/" + endpoint
	}

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		c.AppID+":"+c.AccessToken,
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

func (c *Client) Post(endpoint string, payload []byte) ([]byte, error) {

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api-t1.fyers.in/api/v3/"+endpoint,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		c.AppID+":"+c.AccessToken,
	)

	req.Header.Set("Content-Type", "application/json")

	fmt.Println("===================================")
	fmt.Println("POST URL:", req.URL.String())
	fmt.Println("Authorization:", req.Header.Get("Authorization"))
	fmt.Println("Content-Type:", req.Header.Get("Content-Type"))
	fmt.Println("===================================")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("STATUS =", resp.Status)
	fmt.Println("SERVER =", resp.Header.Get("Server"))
	fmt.Println("CF-RAY =", resp.Header.Get("CF-Ray"))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}
