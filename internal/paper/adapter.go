package paper

import (
	"encoding/json"
	"fmt"

	"github.com/sknain/astracore/broker-go/internal/broker"
)

type Adapter struct {
	engine *Broker
}

func NewAdapter() *Adapter {
	return &Adapter{
		engine: New(),
	}
}

func (a *Adapter) PlaceOrder(req broker.PlaceOrderRequest) (string, error) {

	o, err := a.engine.PlaceOrder(
		req.Symbol,
		Side(req.Side),
		req.Qty,
		req.Price,
	)

	if err != nil {
		return "", err
	}

	return o.ID, nil
}

func (a *Adapter) ModifyOrder(orderID string, qty int, price float64) error {

	return a.engine.ModifyOrder(orderID, qty, price)

}

func (a *Adapter) CancelOrder(orderID string) error {

	return a.engine.CancelOrder(orderID)

}

func (a *Adapter) GetOrder(orderID string) ([]byte, error) {

	a.engine.mu.RLock()
	defer a.engine.mu.RUnlock()

	order, ok := a.engine.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("order not found")
	}

	return json.Marshal(order)
}
