package event

import "sync"

type Subscriber struct {
	ID      string
	Handler Handler
}

type Registry struct {
	mu   sync.RWMutex
	subs map[EventType]map[string]Handler
}

var registry = &Registry{
	subs: make(map[EventType]map[string]Handler),
}

func Register(eventType EventType, id string, handler Handler) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	if registry.subs[eventType] == nil {
		registry.subs[eventType] = make(map[string]Handler)
	}

	registry.subs[eventType][id] = handler

	GetBus().Subscribe(eventType, handler)
}

func Unregister(eventType EventType, id string) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	if registry.subs[eventType] == nil {
		return
	}

	delete(registry.subs[eventType], id)
}
