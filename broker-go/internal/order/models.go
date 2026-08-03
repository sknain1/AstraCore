package order

import "time"

type Status string

const (
	StatusPending   Status = "PENDING"
	StatusOpen      Status = "OPEN"
	StatusFilled    Status = "FILLED"
	StatusCancelled Status = "CANCELLED"
	StatusRejected  Status = "REJECTED"
)

type Side string

const (
	Buy  Side = "BUY"
	Sell Side = "SELL"
)

type Order struct {
	ID string `json:"id"`

	Symbol string `json:"symbol"`

	Side Side `json:"side"`

	Qty int `json:"qty"`

	Price float64 `json:"price"`

	Status Status `json:"status"`

	BrokerID string `json:"broker_id"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}
