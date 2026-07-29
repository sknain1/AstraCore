package broker

import "fmt"

var current Broker

func Register(b Broker) {
	current = b
}

func Get() (Broker, error) {

	if current == nil {
		return nil, fmt.Errorf("no broker registered")
	}

	return current, nil
}
