package workers

import (
	"github.com/sknain/astracore/broker-go/internal/event"
	"github.com/sknain/astracore/broker-go/internal/order"
	"github.com/sknain/astracore/broker-go/internal/portfolio"
)

func RegisterPortfolioWorker() {

	ps := portfolio.NewService()

	event.GetBus().Subscribe(order.OrderFilledEvent, func(e event.Event) {

		o, ok := e.Data.(order.Order)
		if !ok {
			return
		}

		if o.Side == order.Buy {
			_ = ps.ApplyBuy(o.Symbol, o.Qty, o.Price)
		} else {
			_ = ps.ApplySell(o.Symbol, o.Qty, o.Price)
		}
	})
}
