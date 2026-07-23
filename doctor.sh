#!/usr/bin/env bash

set -e

echo "========================================"
echo "      AstraCore Doctor v1.0"
echo "========================================"
echo

PASS=0
FAIL=0
WARN=0

check_ok() {
    echo "✅ $1"
    PASS=$((PASS+1))
}

check_fail() {
    echo "❌ $1"
    FAIL=$((FAIL+1))
}

check_warn() {
    echo "⚠️  $1"
    WARN=$((WARN+1))
}

echo "[1/10] Operating System"

if [ -f /etc/os-release ]; then
    . /etc/os-release
    check_ok "$PRETTY_NAME"
else
    check_fail "OS not detected"
fi

echo
echo "[2/10] Go"

if command -v go >/dev/null 2>&1; then
    check_ok "$(go version)"
else
    check_fail "Go not installed"
fi

echo
echo "[3/10] Rust"

if command -v rustc >/dev/null 2>&1; then
    check_ok "$(rustc --version)"
else
    check_fail "Rust not installed"
fi

echo
echo "[4/10] Cargo"

if command -v cargo >/dev/null 2>&1; then
    check_ok "$(cargo --version)"
else
    check_fail "Cargo not installed"
fi

echo
echo "[5/10] Docker"

if command -v docker >/dev/null 2>&1; then
    check_ok "$(docker --version)"
else
    check_warn "Docker not found"
fi

echo
echo "[6/10] Git"

if command -v git >/dev/null 2>&1; then
    check_ok "$(git --version)"
else
    check_fail "Git not installed"
fi

echo
echo "[7/10] Runtime Folder"

if [ -d "$HOME/.astracore" ]; then
    check_ok "~/.astracore exists"
else
    check_fail "~/.astracore missing"
fi

echo
echo "[8/10] Go Module"

if [ -f broker-go/go.mod ]; then
    check_ok "broker-go/go.mod found"
else
    check_fail "Missing broker-go/go.mod"
fi

echo
echo "[9/10] Rust Project"

if [ -f engine-rust/Cargo.toml ]; then
    check_ok "engine-rust/Cargo.toml found"
else
    check_fail "Missing Cargo.toml"
fi

echo
echo "[10/10] Project"

if [ -d broker-go ] && [ -d engine-rust ]; then
    check_ok "Project structure OK"
else
    check_fail "Project structure invalid"
fi

echo
echo "========================================"
echo "Summary"
echo "========================================"

echo "PASS : $PASS"
echo "WARN : $WARN"
echo "FAIL : $FAIL"

if [ "$FAIL" -eq 0 ]; then
    echo
    echo "🚀 AstraCore is Ready."
else
    echo
    echo "⚠️ Please fix the errors before continuing."
fi
