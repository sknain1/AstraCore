package order

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/sknain/astracore/broker-go/internal/broker"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Place(symbol string, side Side, qty int, price float64) (Order, error) {

	b, err := broker.Get()
	if err != nil {
		return Order{}, err
	}

	brokerSide := broker.Buy
	if side == Sell {
		brokerSide = broker.Sell
	}

	brokerID, err := b.PlaceOrder(broker.PlaceOrderRequest{
		Symbol: symbol,
		Side:   brokerSide,
		Qty:    qty,
		Price:  price,
	})
	if err != nil {
		return Order{}, err
	}

	order := Order{
		ID:        uuid.New().String(),
		BrokerID:  brokerID,
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

	b, err := broker.Get()
	if err != nil {
		return Order{}, err
	}

	if err := b.ModifyOrder(order.BrokerID, qty, price); err != nil {
		return Order{}, err
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

	b, err := broker.Get()
	if err != nil {
		return Order{}, err
	}

	if err := b.CancelOrder(order.BrokerID); err != nil {
		return Order{}, err
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
