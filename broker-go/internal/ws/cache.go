package ws

import "sync"

type TickCache struct {
	mu    sync.RWMutex
	ticks map[string]Tick
}

var cache = &TickCache{
	ticks: make(map[string]Tick),
}

func UpdateTick(tick Tick) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.ticks[tick.Symbol] = tick
}

func GetTick(symbol string) (Tick, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	tick, ok := cache.ticks[symbol]
	return tick, ok
}

func GetAllTicks() []Tick {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	result := make([]Tick, 0, len(cache.ticks))

	for _, tick := range cache.ticks {
		result = append(result, tick)
	}

	return result
}

func ClearTicks() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.ticks = make(map[string]Tick)
}
