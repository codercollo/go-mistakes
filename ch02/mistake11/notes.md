# Mistake 11: Documentation & Linters

(1) Missing Code Documentation
`Rule: Every exported element must be documented`

```go
//BAD — no doc comments
type Customer struct{}
func (c Customer) ID() string { return "" }

//GOOD — starts with element name, complete sentence, ends with punctuation
// Customer is a customer representation.
type Customer struct{}

// ID returns the customer identifier.
func (c Customer) ID() string { return "" }
```

(2) Doc Comment Rules
-Start with the element's name// Customer is a...
-Complete sentence + punctuation// ID returns the customer identifier.
-Describe what, not howWhat it does, not its implementation
-Enough info to use without reading the sourceSelf-contained

(3) Constants & Variables — Two Layers of Comment

```go
// DefaultPermission is the default permission used by the store engine.  ← godoc (purpose)
const DefaultPermission = 0o644 // Need read and write accesses.          ← inline (content)
```

-Doc comment (above) → appears in godoc, for external consumers
-Inline comment (right) → implementation detail, not public-facing

(4) Deprecating Exported Elements

```go
// ComputePath returns the fastest path between two points.
// Deprecated: Uses a legacy algorithm. Use ComputeFastestPath instead.
func ComputePath() string { ... }
```

-IDEs and gopls will show a warning when a deprecated symbol is used.
-Always point to the replacement.

(5) Package Documentation

```go
// Package store provides data persistence operations for customer records.
//
// It supports CRUD operations and handles default permissions
// for all storage engine interactions.
package store
```

Rules:

- Start with // Package <name>- Convention, picked up by godoc
- First line = concise summary
- Extra lines = full detailAfter a blank // line
- File locationSame-name file (store.go) or doc.go
- Must be adjacent to packageA blank line between comment and package breaks it

(6) Common Trap — Non-adjacent Comments Are Ignored

```go
// Copyright 2024 Acme Corp.   ← this will NOT appear in godoc

// Package foo does stuff.     ← this WILL appear (adjacent to package)
package foo
```

# Not Using Linters

Why Linters Matter
`Linters` catch bugs and style issues automatically — things code review misses

shadow detection:

```go
$ go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
$ go vet -vettool=$(which shadow)
# ./main.go:8:3: declaration of "i" shadows declaration at line 6
```

_Essential Linters_

- `go vetBuilt ` -Built-in analyzer — catches common bugs
- `errcheck` -Flags unchecked errors
- `gocyclo` -Flags overly complex functions
- `goconst` - Finds repeated strings that should be constants
- `staticcheck` - Broad static analysis

_Essential Formatters_

- `gofmt` Standard Go Formatter
- `goimports Formats + auto-manages imports

_Use golangci-lint — Run Them All at Once_

- Install
  brew install golangci-lint # or see golangci-lint.run

- Run
  golangci-lint run ./...

Benefits:
Runs many linters in parallel → fast
Single config file per repo
Easy CI integration (GitHub Action included)

_Automate It — Don't Run Manually_
Option 1 — Git pre-commit hook (.git/hooks/pre-commit):
bash#!/bin/sh
golangci-lint run ./...

Option 2 — CI (GitHub Actions) (see .github/lint.yml):
yaml- uses: golangci/golangci-lint-action@v4
