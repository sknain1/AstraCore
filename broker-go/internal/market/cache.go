package market

import "sync"

var (
	mu sync.RWMutex

	indices = make(map[string]Index)
	futures = make(map[string]Future)
	options = make(map[string]Option)
	chains  = make(map[string]OptionChain)
)

func UpdateIndex(idx Index) {
	mu.Lock()
	defer mu.Unlock()

	indices[idx.Name] = idx
}

func GetIndex(name string) (Index, bool) {
	mu.RLock()
	defer mu.RUnlock()

	idx, ok := indices[name]
	return idx, ok
}

func GetAllIndices() map[string]Index {
	mu.RLock()
	defer mu.RUnlock()

	out := make(map[string]Index)
	for k, v := range indices {
		out[k] = v
	}
	return out
}

func UpdateFuture(f Future) {
	mu.Lock()
	defer mu.Unlock()

	futures[f.Symbol] = f
}

func GetFuture(symbol string) (Future, bool) {
	mu.RLock()
	defer mu.RUnlock()

	f, ok := futures[symbol]
	return f, ok
}

func GetAllFutures() map[string]Future {
	mu.RLock()
	defer mu.RUnlock()

	out := make(map[string]Future)
	for k, v := range futures {
		out[k] = v
	}
	return out
}

func UpdateOption(o Option) {
	mu.Lock()
	defer mu.Unlock()

	options[o.Symbol] = o
}

func GetOption(symbol string) (Option, bool) {
	mu.RLock()
	defer mu.RUnlock()

	o, ok := options[symbol]
	return o, ok
}

func GetAllOptions() map[string]Option {
	mu.RLock()
	defer mu.RUnlock()

	out := make(map[string]Option)
	for k, v := range options {
		out[k] = v
	}
	return out
}

func UpdateOptionChain(chain OptionChain) {
	mu.Lock()
	defer mu.Unlock()

	chains[chain.Underlying] = chain
}

func GetOptionChain(symbol string) (OptionChain, bool) {
	mu.RLock()
	defer mu.RUnlock()

	chain, ok := chains[symbol]
	return chain, ok
}

func GetSnapshot() MarketSnapshot {
	mu.RLock()
	defer mu.RUnlock()

	return MarketSnapshot{
		Indices: indices,
		Futures: futures,
		Options: options,
		Chains:  chains,
	}
}
