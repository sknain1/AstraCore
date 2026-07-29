package order

import "sync"

type Manager struct {
	mu sync.RWMutex
}

var (
	manager *Manager
	once    sync.Once
)

func GetManager() *Manager {
	once.Do(func() {
		manager = &Manager{}
	})

	return manager
}

func (m *Manager) Add(order Order) {
	AddOrder(order)
}

func (m *Manager) Update(order Order) {
	UpdateOrder(order)
}

func (m *Manager) Delete(id string) {
	DeleteOrder(id)
}

func (m *Manager) Get(id string) (Order, bool) {
	return GetOrder(id)
}

func (m *Manager) All() []Order {
	return GetOrders()
}

func (m *Manager) Clear() {
	ClearOrders()
}
