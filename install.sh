#!/usr/bin/env bash

set -e

echo "========================================"
echo "      AstraCore Installer v1.0"
echo "========================================"

ROOT="$(pwd)"

echo ""
echo "[1/7] Creating folders..."

mkdir -p \
configs \
docs \
logs \
proto \
scripts \
tests/unit \
tests/integration \
broker-go/cmd \
broker-go/internal/{auth,account,broker,config,grpc,logger,orders,queue,utils,websocket,workers} \
engine-rust/src/{config,grpc,logger,position,risk,strategy,utils}

echo "✓ Folders created"

echo ""
echo "[2/7] Creating root files..."

touch \
README.md \
LICENSE \
.gitignore \
Makefile \
docker-compose.yml \
.env

echo "✓ Root files created"

echo ""
echo "[3/7] Creating config files..."

touch \
configs/app.yaml \
configs/broker.yaml \
configs/risk.yaml \
configs/strategy.yaml \
configs/logging.yaml

echo "✓ Config files created"

echo ""
echo "[4/7] Checking Go..."

if command -v go >/dev/null 2>&1; then
    echo "✓ Go : $(go version)"
else
    echo "✗ Go not installed"
fi

echo ""
echo "[5/7] Checking Rust..."

if command -v cargo >/dev/null 2>&1; then
    echo "✓ Rust : $(rustc --version)"
else
    echo "✗ Rust not installed"
fi

echo ""
echo "[6/7] Checking Docker..."

if command -v docker >/dev/null 2>&1; then
    echo "✓ Docker : $(docker --version)"
else
    echo "⚠ Docker not found"
fi

echo ""
echo "[7/7] Creating AstraCore runtime folders..."

mkdir -p ~/.astracore/{tokens,logs,cache}

echo "✓ Runtime folders created"

echo ""
echo "========================================"
echo " AstraCore Foundation Ready"
echo "========================================"

echo ""
echo "Next Step:"
echo "  ./doctor.sh"
