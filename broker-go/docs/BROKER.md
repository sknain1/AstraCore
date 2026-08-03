# Broker Architecture

## Overview

AstraCore is designed to be broker-agnostic.

The trading engine never communicates directly with a broker.

Instead, every broker implements a common interface.

---

## Architecture

```
                  Order Service
                        │
                        ▼
                Broker Manager
                        │
        ┌───────────────┴───────────────┐
        │                               │
   Paper Adapter                  FYERS Adapter
        │                               │
   Paper Engine                   FYERS REST API
```

Future adapters:

- Zerodha
- Upstox
- Dhan
- Angel One
- Interactive Brokers

---

## Broker Interface

```go
type Broker interface {
    PlaceOrder(req PlaceOrderRequest) (string, error)
    ModifyOrder(orderID string, qty int, price float64) error
    CancelOrder(orderID string) error
    GetOrder(orderID string) ([]byte, error)
}
```

---

## Broker Manager

Responsibilities:

- Register brokers
- Select active broker
- Return current broker
- Support runtime switching

Example:

```go
broker.Register("paper", paper.NewAdapter())
broker.Register("fyers", fyers.NewAdapter(client))

broker.Use("paper")
```

---

## Supported Brokers

| Broker | Status |
|---------|--------|
| Paper | ✅ |
| FYERS | ✅ |
| Zerodha | Planned |
| Upstox | Planned |
| Dhan | Planned |
| Angel One | Planned |
| Interactive Brokers | Planned |

---

## Design Goals

- Broker Independent
- Easy Adapter Integration
- Minimal Code Duplication
- Production Ready
- Extensible
