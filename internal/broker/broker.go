package broker

type Side string

const (
	Buy  Side = "BUY"
	Sell Side = "SELL"
)

type Order struct {
	Symbol string
	Side   Side
	Qty    int
	Price  float64
}

type Broker interface {
	PlaceOrder(order Order) (string, error)
	ModifyOrder(orderID string, qty int, price float64) error
	CancelOrder(orderID string) error
	GetOrder(orderID string) ([]byte, error)
}
