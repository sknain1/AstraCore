package market

import (
	"strings"
	"time"

	"github.com/sknain/astracore/broker-go/internal/event"
	"github.com/sknain/astracore/broker-go/internal/ws"
)

func RegisterSubscribers() {

	event.Register(
		event.TickEvent,
		"market-cache",
		func(e event.Event) {

			tick, ok := e.Data.(ws.Tick)
			if !ok {
				return
			}

			mtick := Tick{
				Symbol:    tick.Symbol,
				LTP:       tick.LTP,
				Timestamp: time.Now(),
			}

			symbol := strings.ToUpper(tick.Symbol)

			// INDIA VIX
			if strings.Contains(symbol, "INDIAVIX") {
				UpdateIndex(Index{
					Name: "INDIA VIX",
					Tick: mtick,
				})
				return
			}

			// NIFTY 50
			if strings.Contains(symbol, "NIFTY50") ||
				strings.Contains(symbol, "NIFTY 50") {
				UpdateIndex(Index{
					Name: "NIFTY 50",
					Tick: mtick,
				})
				return
			}

			// BANKNIFTY
			if strings.Contains(symbol, "BANKNIFTY") {
				UpdateIndex(Index{
					Name: "BANKNIFTY",
					Tick: mtick,
				})
				return
			}

			// Futures
			if strings.Contains(symbol, "FUT") {
				UpdateFuture(Future{
					Symbol: symbol,
					Tick:   mtick,
				})
				return
			}

			// Options
			if strings.Contains(symbol, "CE") ||
				strings.Contains(symbol, "PE") {
				UpdateOption(Option{
					Symbol: symbol,
					Tick:   mtick,
				})
			}
		},
	)
}
