# monotick

Monotonic timestamps as int64 nanoseconds.

## Why int64?

For measuring durations, you don't need wall clock time — just a counter.
int64 nanoseconds can represent ~292 years, while int32 overflows in ~2 seconds.

## Usage

```go
start := monotick.Now()
// ... do work ...
elapsed := monotick.Since(start)
```

## Go version

Requires Go 1.11+ (when `go.mod` was introduced). The code itself uses only
`time` package features available since Go 1.9. Users with modern Go will
build with their version.
