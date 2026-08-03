package market

import "time"

type MarketType string

const (
	IndexType  MarketType = "INDEX"
	FutureType MarketType = "FUTURE"
	OptionType MarketType = "OPTION"
	VIXType    MarketType = "VIX"
)

type Tick struct {
	Symbol string `json:"symbol"`

	Type MarketType `json:"type"`

	LTP float64 `json:"ltp"`

	Open  float64 `json:"open"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
	Close float64 `json:"close"`

	Volume int64 `json:"volume"`

	Bid float64 `json:"bid"`
	Ask float64 `json:"ask"`

	BidQty int64 `json:"bid_qty"`
	AskQty int64 `json:"ask_qty"`

	OI int64 `json:"oi"`

	PrevOI int64 `json:"prev_oi"`

	OIChange int64 `json:"oi_change"`

	IV float64 `json:"iv"`

	Exchange string `json:"exchange"`

	Timestamp time.Time `json:"timestamp"`
}

type Index struct {
	Name string `json:"name"`

	Tick Tick `json:"tick"`
}

type Future struct {
	Symbol string `json:"symbol"`

	Expiry string `json:"expiry"`

	Tick Tick `json:"tick"`
}

type Option struct {
	Symbol string `json:"symbol"`

	Strike float64 `json:"strike"`

	Expiry string `json:"expiry"`

	OptionType string `json:"option_type"`

	Tick Tick `json:"tick"`
}

type OptionChain struct {
	Underlying string `json:"underlying"`

	Expiry string `json:"expiry"`

	Spot float64 `json:"spot"`

	ATM float64 `json:"atm"`

	Options []Option `json:"options"`

	LastUpdated time.Time `json:"last_updated"`
}

type MarketSnapshot struct {
	Indices map[string]Index `json:"indices"`

	Futures map[string]Future `json:"futures"`

	Options map[string]Option `json:"options"`

	Chains map[string]OptionChain `json:"chains"`

	LastUpdated time.Time `json:"last_updated"`
}
