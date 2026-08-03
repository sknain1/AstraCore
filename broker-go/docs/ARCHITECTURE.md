# AstraCore Architecture

## Overview

AstraCore is a modular algorithmic trading operating system written in Go.

The system is broker-agnostic and event-driven.

---

## Core Layers

```
Strategies
      │
      ▼
 Order Service
      │
      ▼
 Broker Manager
      │
 ┌────┴──────────────┐
 │                   │
Paper Broker     FYERS Broker
 │                   │
Local Engine      FYERS REST API
```

---

## Modules

### Strategy Engine

Responsible for generating BUY/SELL signals.

---

### Order Engine

Responsible for:

- Place Orders
- Modify Orders
- Cancel Orders
- Order Cache
- Events

---

### Broker Layer

Provides a unified interface for every broker.

Current brokers:

- Paper
- FYERS

Future brokers:

- Zerodha
- Upstox
- Dhan
- Angel One
- Interactive Brokers

---

### Market Layer

Provides:

- LTP
- Historical Data
- Instruments
- Tick Cache

---

### Risk Engine

Future module.

Responsible for:

- Position sizing
- Daily loss limits
- Exposure limits
- Margin validation

---

### Portfolio Engine

Future module.

Responsible for:

- Holdings
- Positions
- PnL
- Equity Curve

---

## Design Principles

- Broker Agnostic
- Event Driven
- Modular
- Testable
- Clean Architecture
- Production Ready
