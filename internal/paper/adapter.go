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

func (a *Adapter) PlaceOrder(order broker.Order) (string, error) {

	o, err := a.engine.PlaceOrder(
		order.Symbol,
		Side(order.Side),
		order.Qty,
		order.Price,
	)

	if err != nil {
		return "", err
	}

	return o.ID, nil
}

func (a *Adapter) ModifyOrder(orderID string, qty int, price float64) error {

	return fmt.Errorf("paper broker: modify order not implemented")
}

func (a *Adapter) CancelOrder(orderID string) error {

	return fmt.Errorf("paper broker: cancel order not implemented")
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
