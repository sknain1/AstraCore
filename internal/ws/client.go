package ws

import (
	"fmt"
	"sync"
	"time"

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

			fmt.Printf("FYERS MESSAGE: %#v\n", msg)

			symbol, ok := msg["symbol"].(string)
			if !ok {
				return
			}

			var ltp float64

			switch v := msg["ltp"].(type) {
			case float64:
				ltp = v
			case fyersws.FloatSDK:
				ltp = float64(v)
			default:
				fmt.Printf("Unsupported LTP type: %T\n", v)
				return
			}

			tick := Tick{
				Symbol:    symbol,
				LTP:       ltp,
				Timestamp: time.Now(),
			}

			UpdateTick(tick)

			fmt.Println("CACHE UPDATED:", tick.Symbol, tick.LTP)
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
