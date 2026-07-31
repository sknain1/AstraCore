package portfolio

type Position struct {
	Symbol string `json:"symbol"`

	Qty int `json:"qty"`

	AvgPrice float64 `json:"avg_price"`

	LastPrice float64 `json:"last_price"`

	UnrealizedPnL float64 `json:"unrealized_pnl"`

	RealizedPnL float64 `json:"realized_pnl"`
}

type Holding struct {
	Symbol string `json:"symbol"`

	Qty int `json:"qty"`

	AvgPrice float64 `json:"avg_price"`
}

type Portfolio struct {
	Cash float64 `json:"cash"`

	Invested float64 `json:"invested"`

	TotalValue float64 `json:"total_value"`
}
