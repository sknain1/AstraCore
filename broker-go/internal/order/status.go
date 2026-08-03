package order

import "time"

func MarkFilled(id string) (Order, bool) {

	order, ok := GetOrder(id)
	if !ok {
		return Order{}, false
	}

	order.Status = StatusFilled
	order.UpdatedAt = time.Now()

	UpdateOrder(order)

	return order, true
}

func MarkCancelled(id string) (Order, bool) {

	order, ok := GetOrder(id)
	if !ok {
		return Order{}, false
	}

	order.Status = StatusCancelled
	order.UpdatedAt = time.Now()

	UpdateOrder(order)

	return order, true
}

func MarkRejected(id string) (Order, bool) {

	order, ok := GetOrder(id)
	if !ok {
		return Order{}, false
	}

	order.Status = StatusRejected
	order.UpdatedAt = time.Now()

	UpdateOrder(order)

	return order, true
}
