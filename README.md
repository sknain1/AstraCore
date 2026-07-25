# 🚀 AstraCore

AstraCore is a modular algorithmic trading platform built for high-performance, low-latency automated trading.

The project follows a clean architecture where every component is independent, reusable, and production-ready.

---

# 🎯 Vision

Build a complete professional algorithmic trading platform that supports:

- Multiple Brokers
- Live Trading
- Paper Trading
- Strategy Engine
- Risk Management
- Portfolio Management
- Web Dashboard
- Backtesting
- Market Data Streaming

---

# 🛠 Technology Stack

| Component | Technology |
|-----------|------------|
| Broker Layer | Go |
| Trading Engine | Rust (Planned) |
| Communication | gRPC |
| REST APIs | Go net/http |
| Deployment | Docker |
| Version Control | Git + GitHub |

---

# 📂 Project Structure

```text
AstraCore/
│
├── broker-go/
│   ├── cmd/
│   ├── internal/
│   │   ├── auth/
│   │   ├── fyers/
│   │   └── server/
│   │
│   ├── token.json
│   ├── go.mod
│   └── go.sum
│
├── README.md
└── .gitignore
```

---

# ✅ Implemented Features

## Sprint 1

- HTTP Server
- Router
- Health APIs
- Middleware
- Logging
- Recovery Middleware

---

## Sprint 2

### Fyers OAuth

- Login API
- Callback API
- Authorization Code
- Access Token Exchange
- Refresh Token Support
- Token Storage

---

## Sprint 3

### Account APIs

- Profile API
- Funds API
- Holdings API
- Positions API
- Orders API

---

## Sprint 4

### Trading APIs

- Reusable Fyers Client
- Place Order API
- Dry Run Mode
- Modular Client Architecture

---

# 🌐 Available REST APIs

| Method | Endpoint | Status |
|---------|----------|--------|
| GET | /health | ✅ |
| GET | /version | ✅ |
| GET | /ping | ✅ |
| GET | /auth/login | ✅ |
| GET | /auth/callback | ✅ |
| GET | /profile | ✅ |
| GET | /funds | ✅ |
| GET | /holdings | ✅ |
| GET | /positions | ✅ |
| GET | /orders | ✅ |
| POST | /orders/place | ✅ (Dry Run) |

---

# 📌 Current Status

- OAuth Authentication ✅
- Token Management ✅
- Broker Client ✅
- Portfolio APIs ✅
- Trading APIs (Dry Run) ✅

Current Version

```
v0.4.0
```

---

# 🗺 Roadmap

## Sprint 4

- DRY_RUN Configuration
- Live Order Placement
- Modify Order
- Cancel Order
- Trade Book

---

## Sprint 5

- Quotes API
- Market Depth API
- Option Chain API
- Historical Data API

---

## Sprint 6

- WebSocket Market Feed
- Live Tick Processing
- Candle Builder
- Strategy Framework

---

## Sprint 7

- Risk Management
- Position Sizing
- Stop Loss Engine
- Target Engine

---

## Sprint 8

- Portfolio Engine
- PnL Dashboard
- Analytics
- Reports

---

# 🚀 Future Goals

- Multi Broker Support
- Zerodha
- Angel One
- Upstox
- Alice Blue
- Dhan
- Interactive Brokers

---

# 📖 Getting Started

Clone the repository

```bash
git clone https://github.com/sknain1/AstraCore.git
```

Go to project

```bash
cd AstraCore/broker-go
```

Install dependencies

```bash
go mod tidy
```

Run

```bash
go run ./cmd
```

---

# 👨‍💻 Author

**S.K. Nain**

GitHub:

https://github.com/sknain1

---

# 📄 License

This project is under development.

© AstraCore Project
