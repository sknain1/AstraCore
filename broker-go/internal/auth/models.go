package auth

type PlaceOrderRequest struct {
	Symbol       string  `json:"symbol"`
	Qty          int     `json:"qty"`
	Type         int     `json:"type"`
	Side         int     `json:"side"`
	ProductType  string  `json:"productType"`
	LimitPrice   float64 `json:"limitPrice"`
	StopPrice    float64 `json:"stopPrice"`
	DisclosedQty int     `json:"disclosedQty"`
	Validity     string  `json:"validity"`
	OfflineOrder bool    `json:"offlineOrder"`
	StopLoss     float64 `json:"stopLoss"`
	TakeProfit   float64 `json:"takeProfit"`
}
