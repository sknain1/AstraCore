package order

import "fmt"

func Fill(id string) (Order, error) {

	order, ok := MarkFilled(id)
	if !ok {
		return Order{}, fmt.Errorf("order not found")
	}

	PublishFilled(order)

	return order, nil
}
