package fyers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	AppID       string
	AccessToken string
	httpClient  *http.Client
}

func NewClient(appID, accessToken string) *Client {
	return &Client{
		AppID:       appID,
		AccessToken: accessToken,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) doRequest(method, endpoint string, payload []byte) ([]byte, error) {

	url := "https://api-t1.fyers.in/api/v3/" + endpoint

	if len(endpoint) >= 4 && endpoint[:4] == "data" {
		url = "https://api-t1.fyers.in/" + endpoint
	}

	var body io.Reader

	if payload != nil {
		body = bytes.NewBuffer(payload)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		c.AppID+":"+c.AccessToken,
	)
	req.Header.Set("Content-Type", "application/json")

	fmt.Println("========================================")
	fmt.Println("FYERS REQUEST")
	fmt.Println("METHOD :", method)
	fmt.Println("URL    :", url)
	fmt.Println("APP ID :", c.AppID)

	auth := c.AppID + ":" + c.AccessToken
	if len(auth) > 35 {
		fmt.Println("AUTH   :", auth[:35]+"...")
	} else {
		fmt.Println("AUTH   :", auth)
	}

	if payload != nil {
		fmt.Println("BODY   :", string(payload))
	}
	fmt.Println("========================================")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("HTTP ERROR:", err)
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("========================================")
	fmt.Println("FYERS RESPONSE")
	fmt.Println("STATUS :", resp.Status)
	fmt.Println("BODY   :", string(respBody))
	fmt.Println("========================================")

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", respBody)
	}

	return respBody, nil
}

func (c *Client) Get(endpoint string) ([]byte, error) {
	return c.doRequest(http.MethodGet, endpoint, nil)
}

func (c *Client) Post(endpoint string, payload []byte) ([]byte, error) {
	return c.doRequest(http.MethodPost, endpoint, payload)
}

func (c *Client) Patch(endpoint string, payload []byte) ([]byte, error) {
	return c.doRequest(http.MethodPatch, endpoint, payload)
}

func (c *Client) Delete(endpoint string, payload []byte) ([]byte, error) {
	return c.doRequest(http.MethodDelete, endpoint, payload)
}
