package ws

import "time"

type Tick struct {
	Symbol    string    `json:"symbol"`
	LTP       float64   `json:"ltp"`
	Bid       float64   `json:"bid"`
	Ask       float64   `json:"ask"`
	Open      float64   `json:"open"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Close     float64   `json:"close"`
	Volume    int64     `json:"volume"`
	Timestamp time.Time `json:"timestamp"`
}

type SubscribeRequest struct {
	Symbols []string `json:"symbols"`
}
