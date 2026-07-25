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

	req, err := http.NewRequest(
		http.MethodGet,
		"https://api-t1.fyers.in/api/v3/"+endpoint,
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
