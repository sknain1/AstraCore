package paper

func (b *Broker) Funds() Funds {

	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.funds
}
