package order

import "sync"

type Cache struct {
	mu     sync.RWMutex
	orders map[string]Order
}

var cache = &Cache{
	orders: make(map[string]Order),
}

func AddOrder(order Order) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.orders[order.ID] = order
}

func UpdateOrder(order Order) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.orders[order.ID] = order
}

func GetOrder(id string) (Order, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	order, ok := cache.orders[id]
	return order, ok
}

func GetOrders() []Order {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	orders := make([]Order, 0, len(cache.orders))

	for _, order := range cache.orders {
		orders = append(orders, order)
	}

	return orders
}

func DeleteOrder(id string) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	delete(cache.orders, id)
}

func ClearOrders() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.orders = make(map[string]Order)
}
