package order

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Place(symbol string, side Side, qty int, price float64) (Order, error) {

	order := Order{
		ID:        uuid.New().String(),
		Symbol:    symbol,
		Side:      side,
		Qty:       qty,
		Price:     price,
		Status:    StatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	AddOrder(order)
	PublishPlaced(order)

	return order, nil
}

func (s *Service) Modify(id string, qty int, price float64) (Order, error) {

	order, ok := GetOrder(id)
	if !ok {
		return Order{}, fmt.Errorf("order not found")
	}

	order.Qty = qty
	order.Price = price
	order.UpdatedAt = time.Now()

	UpdateOrder(order)
	PublishModified(order)

	return order, nil
}

func (s *Service) Cancel(id string) (Order, error) {

	order, ok := GetOrder(id)
	if !ok {
		return Order{}, fmt.Errorf("order not found")
	}

	order.Status = StatusCancelled
	order.UpdatedAt = time.Now()

	UpdateOrder(order)
	PublishCancelled(order)

	return order, nil
}

func (s *Service) Get(id string) (Order, bool) {
	return GetOrder(id)
}

func (s *Service) All() []Order {
	return GetOrders()
}
