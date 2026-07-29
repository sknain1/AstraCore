package ws

import (
	"fmt"
	"sync"

	fyersws "github.com/FyersDev/fyers-go-sdk/websocket"
)

type Client struct {
	socket *fyersws.FyersDataSocket

	mu sync.RWMutex

	appID       string
	accessToken string

	connected bool
}

func NewClient(appID, accessToken string) *Client {
	return &Client{
		appID:       appID,
		accessToken: accessToken,
	}
}

func (c *Client) Connect() error {

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.socket != nil && c.socket.IsConnected() {
		c.connected = true
		return nil
	}

	accessToken := c.appID + ":" + c.accessToken

	c.socket = fyersws.NewFyersDataSocket(
		accessToken,
		"",
		true,
		false,
		true,
		5,
		func() {
			fmt.Println("WebSocket Connected")
			c.connected = true
		},
		func(msg fyersws.DataClose) {
			fmt.Println("WebSocket Closed:", msg)
			c.connected = false
		},
		func(err fyersws.DataError) {
			fmt.Println("WebSocket Error:", err)
			c.connected = false
		},
		func(msg fyersws.DataResponse) {
			fmt.Println(msg)
		},
	)

	if c.socket == nil {
		return fmt.Errorf("failed to create websocket client")
	}

	return c.socket.Connect()
}

func (c *Client) Disconnect() {

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.socket != nil {
		c.socket.CloseConnection()
		c.connected = false
	}
}

func (c *Client) Subscribe(symbols []string, dataType string) error {

	if c.socket == nil {
		return fmt.Errorf("websocket not connected")
	}

	c.socket.Subscribe(symbols, dataType)

	return nil
}

func (c *Client) IsConnected() bool {

	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.socket == nil {
		return false
	}

	return c.socket.IsConnected()
}
