package ws

import "sync"

type TickCache struct {
	mu    sync.RWMutex
	ticks map[string]Tick
}

func NewTickCache() *TickCache {
	return &TickCache{
		ticks: make(map[string]Tick),
	}
}

func (c *TickCache) Set(t Tick) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.ticks[t.Symbol] = t
}

func (c *TickCache) Get(symbol string) (Tick, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	t, ok := c.ticks[symbol]
	return t, ok
}

func (c *TickCache) GetAll() []Tick {
	c.mu.RLock()
	defer c.mu.RUnlock()

	out := make([]Tick, 0, len(c.ticks))

	for _, t := range c.ticks {
		out = append(out, t)
	}

	return out
}
