package paper

import "sync"

type Broker struct {
	mu sync.RWMutex

	orders map[string]Order

	positions map[string]Position

	holdings map[string]Holding

	lastTick map[string]Tick

	funds Funds
}

func New() *Broker {

	return &Broker{

		orders: make(map[string]Order),

		positions: make(map[string]Position),

		holdings: make(map[string]Holding),

		lastTick: make(map[string]Tick),

		funds: Funds{
			Balance: 1000000,
			Used:    0,
			Free:    1000000,
		},
	}
}
