# AstraCore

> A Broker-Agnostic Algorithmic Trading Operating System written in Go.

---

# Vision

AstraCore is designed to become a production-grade algorithmic trading platform.

It provides a clean architecture for:

- Multi Broker Trading
- Strategy Engine
- Risk Engine
- Portfolio Management
- Live Trading
- Backtesting
- Event Driven Execution

---

# Current Features

## Broker Layer

- ✅ Paper Broker
- ✅ FYERS Broker
- ✅ Multi Broker Manager
- ✅ Broker Agnostic Order Engine

---

## Market Layer

- Instrument Loader
- Tick Cache
- WebSocket Manager
- Market Subscribers

---

## Order Engine

- Place Order
- Modify Order
- Cancel Order
- Order Cache
- Event Publishing

---

## Infrastructure

- REST API
- WebSocket
- Event Bus
- Config Loader
- Logger

---

# Project Structure

```
broker-go/

cmd/
configs/
docs/
internal/

    account/
    auth/
    broker/
    config/
    event/
    fyers/
    grpc/
    instruments/
    logger/
    market/
    order/
    paper/
    queue/
    server/
    utils/
    workers/
    ws/
```

---

# Build

```
go fmt ./...
go build ./...
go test ./...
```

---

# Documentation

See:

- docs/ARCHITECTURE.md
- docs/ROADMAP.md
- docs/BROKER.md

---

# Roadmap

Current Version:

**v0.3.x**

Upcoming:

- Paper Broker Completion
- Portfolio Engine
- Risk Engine
- Strategy Engine
- Backtesting
- Zerodha Adapter
- Upstox Adapter
- Live Trading

---

# Design Principles

- Clean Architecture
- Broker Agnostic
- Event Driven
- Modular
- Production Ready
- Extensible

---

# License

Private Project

Copyright © AstraCore
