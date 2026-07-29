package order

import (
	"time"

	"github.com/sknain/astracore/broker-go/internal/event"
)

const (
	OrderPlacedEvent    event.EventType = "ORDER_PLACED"
	OrderModifiedEvent  event.EventType = "ORDER_MODIFIED"
	OrderCancelledEvent event.EventType = "ORDER_CANCELLED"
	OrderFilledEvent    event.EventType = "ORDER_FILLED"
)

func PublishPlaced(order Order) {
	event.GetBus().Publish(event.Event{
		Type:      OrderPlacedEvent,
		Timestamp: time.Now(),
		Data:      order,
	})
}

func PublishModified(order Order) {
	event.GetBus().Publish(event.Event{
		Type:      OrderModifiedEvent,
		Timestamp: time.Now(),
		Data:      order,
	})
}

func PublishCancelled(order Order) {
	event.GetBus().Publish(event.Event{
		Type:      OrderCancelledEvent,
		Timestamp: time.Now(),
		Data:      order,
	})
}

func PublishFilled(order Order) {
	event.GetBus().Publish(event.Event{
		Type:      OrderFilledEvent,
		Timestamp: time.Now(),
		Data:      order,
	})
}
