package order

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/sknain/astracore/broker-go/internal/auth"
	"github.com/sknain/astracore/broker-go/internal/fyers"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Place(symbol string, side Side, qty int, price float64) (Order, error) {

	client, err := auth.NewFyersClient()
	if err != nil {
		return Order{}, err
	}

	fyersSide := 1
	if side == Sell {
		fyersSide = -1
	}

	resp, err := client.PlaceOrder(fyers.PlaceOrderRequest{
		Symbol:       symbol,
		Qty:          qty,
		Type:         1, // Limit Order
		Side:         fyersSide,
		ProductType:  "INTRADAY",
		LimitPrice:   price,
		StopPrice:    0,
		DisclosedQty: 0,
		Validity:     "DAY",
		OfflineOrder: false,
	})
	if err != nil {
		return Order{}, err
	}

	var brokerResp map[string]any
	_ = json.Unmarshal(resp, &brokerResp)

	brokerID := ""
	if id, ok := brokerResp["id"].(string); ok {
		brokerID = id
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

	client, err := auth.NewFyersClient()
	if err != nil {
		return Order{}, err
	}

	_, err = client.ModifyOrder(fyers.ModifyOrderRequest{
		ID:         order.BrokerID,
		Qty:        qty,
		LimitPrice: price,
		Type:       1,
	})
	if err != nil {
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

	client, err := auth.NewFyersClient()
	if err != nil {
		return Order{}, err
	}

	_, err = client.CancelOrder(order.BrokerID)
	if err != nil {
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
