package portfolio

import (
	"fmt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Portfolio() Portfolio {

	mu.RLock()
	defer mu.RUnlock()

	return account
}

func (s *Service) Positions() []Position {

	mu.RLock()
	defer mu.RUnlock()

	out := make([]Position, 0, len(positions))

	for _, p := range positions {
		out = append(out, p)
	}

	return out
}

func (s *Service) Holdings() []Holding {

	mu.RLock()
	defer mu.RUnlock()

	out := make([]Holding, 0, len(holdings))

	for _, h := range holdings {
		out = append(out, h)
	}

	return out
}

func (s *Service) ApplyBuy(symbol string, qty int, price float64) error {

	mu.Lock()
	defer mu.Unlock()

	account.Cash -= float64(qty) * price
	account.Invested += float64(qty) * price

	pos := positions[symbol]

	if pos.Qty == 0 {

		pos.Symbol = symbol
		pos.Qty = qty
		pos.AvgPrice = price
		pos.LastPrice = price

	} else {

		totalQty := pos.Qty + qty
		pos.AvgPrice = ((pos.AvgPrice * float64(pos.Qty)) +
			(price * float64(qty))) / float64(totalQty)

		pos.Qty = totalQty
		pos.LastPrice = price
	}

	positions[symbol] = pos

	holdings[symbol] = Holding{
		Symbol:   symbol,
		Qty:      pos.Qty,
		AvgPrice: pos.AvgPrice,
	}

	account.TotalValue = account.Cash + account.Invested

	return nil
}

func (s *Service) ApplySell(symbol string, qty int, price float64) error {

	mu.Lock()
	defer mu.Unlock()

	pos, ok := positions[symbol]
	if !ok {
		return fmt.Errorf("position not found")
	}

	if qty > pos.Qty {
		return fmt.Errorf("insufficient quantity")
	}

	cost := pos.AvgPrice * float64(qty)
	proceeds := price * float64(qty)

	account.Cash += proceeds
	account.Invested -= cost

	pos.Qty -= qty
	pos.LastPrice = price
	pos.RealizedPnL += proceeds - cost

	if pos.Qty == 0 {
		delete(positions, symbol)
		delete(holdings, symbol)
	} else {
		positions[symbol] = pos

		holdings[symbol] = Holding{
			Symbol:   symbol,
			Qty:      pos.Qty,
			AvgPrice: pos.AvgPrice,
		}
	}

	account.TotalValue = account.Cash + account.Invested

	return nil
}
