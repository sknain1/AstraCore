package ws

import (
	"fmt"

	"github.com/sknain/astracore/broker-go/internal/event"
)

func RegisterSubscribers() {

	fmt.Println("Registering Tick Subscriber...")

	event.Register(
		event.TickEvent,
		"tick-cache",
		func(e event.Event) {

			fmt.Println("Tick Event Received")

			tick, ok := e.Data.(Tick)
			if !ok {
				fmt.Println("Type assertion failed")
				return
			}

			UpdateTick(tick)

			fmt.Println("Tick Cached:", tick.Symbol)
		},
	)
}
