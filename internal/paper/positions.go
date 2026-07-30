package paper

func (b *Broker) Positions() []Position {

	b.mu.RLock()
	defer b.mu.RUnlock()

	list := make([]Position, 0, len(b.positions))

	for _, p := range b.positions {
		list = append(list, p)
	}

	return list
}

func (b *Broker) Position(symbol string) (Position, bool) {

	b.mu.RLock()
	defer b.mu.RUnlock()

	p, ok := b.positions[symbol]

	return p, ok
}

func (b *Broker) updatePosition(symbol string, qty int, price float64, side Side) {

	p := b.positions[symbol]

	if side == Buy {

		totalQty := p.Qty + qty

		if totalQty > 0 {
			p.Avg = ((p.Avg * float64(p.Qty)) + (price * float64(qty))) / float64(totalQty)
		}

		p.Qty = totalQty

	} else {

		p.Qty -= qty

		if p.Qty <= 0 {
			delete(b.positions, symbol)
			return
		}
	}

	b.positions[symbol] = p
}
