package portfolio

import "sync"

var (
	mu sync.RWMutex

	positions = make(map[string]Position)

	holdings = make(map[string]Holding)

	account = Portfolio{
		Cash:       1000000,
		Invested:   0,
		TotalValue: 1000000,
	}
)
