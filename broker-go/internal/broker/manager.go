package broker

import "fmt"

var (
	brokers = make(map[string]Broker)
	current Broker
)

func Register(name string, b Broker) {
	brokers[name] = b
}

func Use(name string) error {

	b, ok := brokers[name]
	if !ok {
		return fmt.Errorf("broker '%s' not found", name)
	}

	current = b
	return nil
}

func Get() (Broker, error) {

	if current == nil {
		return nil, fmt.Errorf("no broker selected")
	}

	return current, nil
}

func List() []string {

	out := make([]string, 0, len(brokers))

	for name := range brokers {
		out = append(out, name)
	}

	return out
}
