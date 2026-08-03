package paper

func (b *Broker) Holdings() []Holding {

	b.mu.RLock()
	defer b.mu.RUnlock()

	list := make([]Holding, 0, len(b.holdings))

	for _, h := range b.holdings {
		list = append(list, h)
	}

	return list
}

func (b *Broker) Holding(symbol string) (Holding, bool) {

	b.mu.RLock()
	defer b.mu.RUnlock()

	h, ok := b.holdings[symbol]

	return h, ok
}
