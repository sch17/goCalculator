# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A command-line calculator (written in Go) that performs arithmetic on either Arabic or Roman numerals. User input and error messages are in Russian. There is no `go.mod` — this is a single standalone file with no external dependencies.

## Commands

```bash
# Run the calculator interactively
go run main.go

# Lint/vet
go vet main.go

# Build binary
go build -o calculator main.go

# Run tests (once test files exist)
go test ./...

# Run a single test by name
go test -run TestCalc
```

Input format is `<operand1> <operator> <operand2>` (space-separated), e.g. `3 + 4` or `X - IX`.

## Architecture

Everything lives in `main.go` (package `main`). The call flow is:

```
main() → REPL loop
  └─ calc(s string) → parses the 3-token expression, type-checks operands
       ├─ isDecimalOrRoman()   — determines "decimal" or "roman" for each operand
       ├─ convertToDecimal()   — dispatches to strconv.Atoi or getDecFromRom()
       ├─ action()             — performs +, -, *, /
       └─ getRomFromDec()      — converts result back to Roman if needed
```

Key constraints enforced by `calc()`:
- Both operands must be the same type (cannot mix Arabic and Roman).
- Roman numeral inputs are only recognized for `I`–`X` (1–10); anything else returns 0 from `getDecFromRom`.
- Roman numeral results must be ≥ 1 (no zero or negative Roman numbers).
- `getRomFromDec` handles results up to 100 (C).

On any error, `main()` prints the error and **returns** (exits the loop), so the program terminates rather than continuing.

## Adding Tests

There is no test file yet. To add one, create `main_test.go` in the same directory with `package main` and call `calc()` / `action()` directly — all functions are package-level and accessible within the same package.
