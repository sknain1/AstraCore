package paper

type Side string

const (
	Buy  Side = "BUY"
	Sell Side = "SELL"
)

type OrderStatus string

const (
	OrderPending   OrderStatus = "PENDING"
	OrderFilled    OrderStatus = "FILLED"
	OrderCancelled OrderStatus = "CANCELLED"
	OrderRejected  OrderStatus = "REJECTED"
)

type Order struct {
	ID        string
	Symbol    string
	Side      Side
	Qty       int
	Price     float64
	Status    OrderStatus
	Timestamp int64
}

type Position struct {
	Symbol string
	Qty    int
	Avg    float64
	LTP    float64
	PnL    float64
}

type Holding struct {
	Symbol string
	Qty    int
	Avg    float64
	LTP    float64
}

type Funds struct {
	Balance float64
	Used    float64
	Free    float64
}

type Tick struct {
	Symbol string
	LTP    float64
}
