package event

import "sync"

type Handler func(Event)

type Bus struct {
	mu          sync.RWMutex
	subscribers map[EventType][]Handler
}

var (
	bus  *Bus
	once sync.Once
)

func GetBus() *Bus {
	once.Do(func() {
		bus = &Bus{
			subscribers: make(map[EventType][]Handler),
		}
	})
	return bus
}

func (b *Bus) Subscribe(eventType EventType, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.subscribers[eventType] = append(b.subscribers[eventType], handler)
}

func (b *Bus) Publish(event Event) {
	b.mu.RLock()
	handlers := append([]Handler{}, b.subscribers[event.Type]...)
	b.mu.RUnlock()

	for _, h := range handlers {
		go h(event)
	}
}
