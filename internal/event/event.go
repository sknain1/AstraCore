package event

import "time"

// EventType represents different event categories.
type EventType string

const (
	TickEvent     EventType = "tick"
	OrderEvent    EventType = "order"
	PositionEvent EventType = "position"
	HoldingEvent  EventType = "holding"
	RiskEvent     EventType = "risk"
	SystemEvent   EventType = "system"
)

// Event is the base event passed through the Event Bus.
type Event struct {
	Type      EventType
	Timestamp time.Time
	Data      any
}
