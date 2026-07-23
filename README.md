# AstraCore

## 🚀 Overview

AstraCore is a modular trading system architecture designed for high-performance algorithmic trading.

Technology Stack:

- Go for Broker Connectivity
- Rust for Trading Engine
- gRPC for Communication
- Docker for Deployment

---

## 🎯 Vision

A reliable, fast and expandable trading platform where:

- Multiple brokers can be connected
- Trading strategies can be added easily
- Risk management stays separated
- Trading engine runs independently

---

## 🏗 Architecture

            AstraCore

                |
                |
          +-------------+
          |  Go Layer   |
          | Broker API  |
          | WebSocket   |
          +-------------+

                |
              gRPC

                |

          +-------------+
          | Rust Engine |
          | Strategy    |
          | Risk        |
          | Position    |
          +-------------+

---

## 🧰 Technology Stack

### Backend

- Go 1.23
- Rust
- Docker
- gRPC

### Broker Layer

- Fyers API
- REST API
- WebSocket

### Future Components

- Redis
- Monitoring
- Dashboard

---

## 📁 Project Structure


---

## ⚙ Installation

Clone project:

```bash
git clone <repository>
cd AstraCore
./install.sh
./doctor.sh
make build
make test
gofmt -w .
cargo fmt
~/.astracore/tokens/
.env

चलाने के बाद जाँच करें:

```bash
cat README.md
git add README.md
git commit -m "Sprint-1 D4: README documentation"
