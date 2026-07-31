# Contributing to AstraCore

## Coding Standards

- Always run:

```bash
go fmt ./...
go build ./...
go test ./...
```

before committing.

---

## Commit Style

Examples:

```
feat(order): add order cache

fix(ws): reconnect bug

refactor(broker): broker agnostic order service

docs: update roadmap
```

---

## Branch Strategy

master

↓

feature branches

↓

merge

---

## Pull Request Checklist

- Build passes
- Tests pass
- Documentation updated
- Clean commit history

---

## Design Principles

- Clean Architecture
- Modular
- Broker Agnostic
- Event Driven
- Production Ready
