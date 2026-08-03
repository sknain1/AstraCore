package paper

import (
	"fmt"
	"time"
)

func (b *Broker) PlaceOrder(
	symbol string,
	side Side,
	qty int,
	price float64,
) (Order, error) {

	b.mu.Lock()
	defer b.mu.Unlock()

	id := fmt.Sprintf("%d", time.Now().UnixNano())

	order := Order{
		ID:        id,
		Symbol:    symbol,
		Side:      side,
		Qty:       qty,
		Price:     price,
		Status:    OrderFilled,
		Timestamp: time.Now().Unix(),
	}

	b.orders[id] = order

	b.updatePosition(symbol, qty, price, side)

	return order, nil
}

func (b *Broker) Orders() []Order {

	b.mu.RLock()
	defer b.mu.RUnlock()

	orders := make([]Order, 0, len(b.orders))

	for _, o := range b.orders {
		orders = append(orders, o)
	}

	return orders
}

func (b *Broker) ModifyOrder(id string, qty int, price float64) error {

	b.mu.Lock()
	defer b.mu.Unlock()

	order, ok := b.orders[id]
	if !ok {
		return fmt.Errorf("order not found")
	}

	order.Qty = qty
	order.Price = price

	b.orders[id] = order

	return nil
}

func (b *Broker) CancelOrder(id string) error {

	b.mu.Lock()
	defer b.mu.Unlock()

	order, ok := b.orders[id]
	if !ok {
		return fmt.Errorf("order not found")
	}

	order.Status = OrderCancelled

	b.orders[id] = order

	return nil
}
